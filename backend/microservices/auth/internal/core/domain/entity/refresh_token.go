package entity

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

type RefreshToken struct {
	Token   string `json:"refresh_token"`
	OwnerId string `json:"id"`
}

func HashToken(refreshToken string) (string, error) {

	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(refreshToken))
	sha256Token := sha256Hash.Sum(nil)

	sha256TokenHex := hex.EncodeToString(sha256Token)

	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(sha256TokenHex), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptHash), nil
}

func CheckTokenHash(plainToken, hashedToken string) bool {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(plainToken))
	sha256Token := sha256Hash.Sum(nil)

	sha256TokenHex := hex.EncodeToString(sha256Token)

	err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(sha256TokenHex))
	return err == nil
}
