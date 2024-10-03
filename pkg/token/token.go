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
	CreateTokenString(data any) (string, error)
	ExtractTokenData(tokenString string, data any) error
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

type Payload struct {
	Data []byte `json:"data"`
	jwt.RegisteredClaims
}

func (token *token) CreateTokenString(data any) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		errStr := fmt.Sprintf("error marshal data: %v", err)
		return "", errors.New(errStr)
	}

	expiredAt := jwt.NewNumericDate(time.Now().Add(token.accessTokenExpiration))
	registeredClaim := jwt.RegisteredClaims{ExpiresAt: expiredAt}
	payload := &Payload{dataBytes, registeredClaim}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, payload)
	return jwtToken.SignedString(token.privateEd25519Key)
}

func (token *token) CreateRefreshTokenString(data any) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		errStr := fmt.Sprintf("error marshal data: %v", err)
		return "", errors.New(errStr)
	}

	expiredAt := jwt.NewNumericDate(time.Now().Add(token.refreshTokenExpiration))
	registeredClaim := jwt.RegisteredClaims{
		ExpiresAt: expiredAt,
		ID:        "refresh", // Unique claim to distinguish refresh token
	}
	payload := &Payload{dataBytes, registeredClaim}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, payload)
	return jwtToken.SignedString(token.privateEd25519Key)
}

const (
	inValidToken        = "invalid token"
	errorMappingPayload = "error mapping the payload"
	errorUnmarshalData  = "error unmarshaling the data"
)

func (token *token) ExtractTokenData(tokenString string, data any) error {
	checkSigningMethod := func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("wrong signing method: %v", jwtToken.Header["alg"])
		}
		return token.publicEd25519Key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &Payload{}, checkSigningMethod, jwt.WithoutClaimsValidation())
	if err != nil {
		errStr := fmt.Sprintf("error: %v, token: %s", err, tokenString)
		return errors.New(errStr)
	}

	if !jwtToken.Valid {
		errStr := fmt.Sprintf("%s, token: %v", "Invalid token", jwtToken)
		return errors.New(errStr)
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		errStr := fmt.Sprintf("%s: %s, token: %v", "Invalid token", "error mapping payload", jwtToken)
		return errors.New(errStr)
	}

	if err := json.Unmarshal([]byte(payload.Data), data); err != nil {
		errStr := fmt.Sprintf("%s: %s, data: %s", "Invalid token", "error unmarshalling data", payload.Data)
		return errors.New(errStr)
	}

	if payload.ExpiresAt != nil && time.Now().After(payload.ExpiresAt.Time) {
		return errors.New("error token has expired")
	}

	fmt.Println(payload.ExpiresAt)
	fmt.Println(payload.ExpiresAt.Time)
	fmt.Println(time.Now())
	fmt.Println(time.Now().After(payload.ExpiresAt.Time))

	return nil
}

func (token *token) ValidateRefreshToken(tokenString string) error {
	checkSigningMethod := func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("wrong signing method: %v", jwtToken.Header["alg"])
		}
		return token.publicEd25519Key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &Payload{}, checkSigningMethod)
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
