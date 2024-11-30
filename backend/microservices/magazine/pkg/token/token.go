package token

import (
	"crypto"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	ExtractTokenData(tokenString string) (*AccessTokenPayload, error)
}

type token struct {
	publicEd25519Key      crypto.PublicKey
	accessTokenExpiration time.Duration
}

func New(cfg *Config) (Token, error) {
	token := &token{}
	var err error

	publicPemKey := []byte(cfg.PublicPem)
	token.publicEd25519Key, err = jwt.ParseEdPublicKeyFromPEM(publicPemKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Ed25519 public key: %v", err)
	}

	token.accessTokenExpiration = cfg.AccessTokenExpiration

	return token, nil
}

type AccessTokenPayload struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

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
