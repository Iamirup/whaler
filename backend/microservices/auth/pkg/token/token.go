package token

import (
	"crypto"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	CreateTokenString(userId, username string, isAdmin bool) (string, error)
	ExtractTokenData(tokenString string) (*AccessTokenPayload, error)
	CreateRefreshTokenString(data any) (string, error)
	ValidateRefreshToken(tokenString string) error
	GetRefreshTokenExpiration() time.Duration
}

type token struct {
	privateEd25519Key      crypto.PrivateKey
	publicEd25519Key       crypto.PublicKey
	accessTokenExpiration  time.Duration
	refreshTokenExpiration time.Duration
}

func New(cfg *Config) (Token, error) {
	token := &token{}
	var err error

	privatePemKey := []byte(cfg.PrivatePem)
	token.privateEd25519Key, err = jwt.ParseEdPrivateKeyFromPEM(privatePemKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Ed25519 private key: %v", err)
	}

	publicPemKey := []byte(cfg.PublicPem)
	token.publicEd25519Key, err = jwt.ParseEdPublicKeyFromPEM(publicPemKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Ed25519 public key: %v", err)
	}

	token.accessTokenExpiration = cfg.AccessTokenExpiration
	token.refreshTokenExpiration = cfg.RefreshTokenExpiration

	return token, nil
}

type AccessTokenPayload struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

type RefreshTokenPayload struct {
	Data []byte `json:"data"`
	jwt.RegisteredClaims
}

func (token *token) CreateTokenString(userId, username string, isAdmin bool) (string, error) {
	payload := &AccessTokenPayload{
		Id:       userId,
		Username: username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(token.accessTokenExpiration)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, payload)
	return jwtToken.SignedString(token.privateEd25519Key)
}

func (token *token) CreateRefreshTokenString(data any) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		errStr := fmt.Sprintf("error marshal data: %v", err)
		return "", errors.New(errStr)
	}

	registeredClaim := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(token.refreshTokenExpiration)),
		// ID:        fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Microsecond)),
	}
	payload := &RefreshTokenPayload{dataBytes, registeredClaim}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, payload)
	return jwtToken.SignedString(token.privateEd25519Key)
}

const (
	inValidToken        = "invalid token"
	errorMappingPayload = "error mapping the payload"
	errorUnmarshalData  = "error unmarshaling the data"
)

func (token *token) ExtractTokenData(tokenString string) (*AccessTokenPayload, error) {
	checkSigningMethod := func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("wrong signing method: %v", jwtToken.Header["alg"])
		}
		return token.publicEd25519Key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &AccessTokenPayload{}, checkSigningMethod, jwt.WithoutClaimsValidation())
	if err != nil {
		errStr := fmt.Sprintf("error: %v, token: %s", err, tokenString)
		return nil, errors.New(errStr)
	}

	if !jwtToken.Valid {
		errStr := fmt.Sprintf("%s, token: %v", "Invalid token", jwtToken)
		return nil, errors.New(errStr)
	}

	payload, ok := jwtToken.Claims.(*AccessTokenPayload)
	if !ok {
		errStr := fmt.Sprintf("%s: %s, token: %v", "Invalid token", "error mapping payload", jwtToken)
		return nil, errors.New(errStr)
	}

	if payload.ExpiresAt != nil && time.Now().After(payload.ExpiresAt.Time) {
		return payload, errors.New("error token has expired")
	}

	return payload, nil
}

func (token *token) ValidateRefreshToken(tokenString string) error {
	checkSigningMethod := func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("wrong signing method: %v", jwtToken.Header["alg"])
		}
		return token.publicEd25519Key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &RefreshTokenPayload{}, checkSigningMethod)
	if err != nil {
		errStr := fmt.Sprintf("error: %v, token: %s", err, tokenString)
		return errors.New(errStr)
	}

	if !jwtToken.Valid {
		errStr := fmt.Sprintf("%s, token: %v", inValidToken, jwtToken)
		return errors.New(errStr)
	}

	return nil
}

func (token *token) GetRefreshTokenExpiration() time.Duration {
	return token.refreshTokenExpiration
}
