package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id         string     `json:"Id"`
	Username   string     `json:"username"`
	Password   string     `json:"password,omitempty"`
	UserConfig UserConfig `json:"user_config,omitempty"`
	CreatedAt  string     `json:"created_at"`
}

func (c User) Marshal() *User {
	return &User{
		Id:         c.Id,
		Username:   c.Username,
		UserConfig: c.UserConfig,
		CreatedAt:  c.CreatedAt,
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
