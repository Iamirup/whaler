package entity

import "golang.org/x/crypto/bcrypt"

type RefreshToken struct {
	Token   string `json:"refresh_token"`
	OwnerId string `json:"id"`
}

func HashToken(refreshToken string) (string, error) {
	tokenBytes := []byte(refreshToken)
	hashedTokenBytes, err := bcrypt.GenerateFromPassword(tokenBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedTokenBytes), nil
}

func CheckTokenHash(plainToken, hashedToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(plainToken))
	return err == nil
}
