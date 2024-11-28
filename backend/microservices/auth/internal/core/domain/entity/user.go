package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (c User) Marshal() *User {
	return &User{
		Id:        c.Id,
		Username:  c.Username,
		CreatedAt: c.CreatedAt,
	}
}

func (u *User) HashPassword() (string, error) {
	passwordBytes := []byte(u.Password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

func (u *User) CheckPasswordHash(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err == nil
}
