package entity

import (
	"crypto/sha256"
	"encoding/hex"
)

type RefreshToken struct {
	Token   string `json:"refresh_token"`
	OwnerId string `json:"id"`
}

func HashToken(refreshToken, pepper string) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(refreshToken + pepper))
	sha256Token := sha256Hash.Sum(nil)

	sha256TokenHex := hex.EncodeToString(sha256Token)

	return sha256TokenHex
}

func CheckTokenHash(plainToken, hashedToken, pepper string) bool {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(plainToken + pepper))
	sha256Token := sha256Hash.Sum(nil)

	sha256TokenHex := hex.EncodeToString(sha256Token)

	return sha256TokenHex == plainToken
}
