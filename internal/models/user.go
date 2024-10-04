package models

import (
	"github.com/Iamirup/whaler/pkg/utils"
)

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

func (c User) HashPassword() (string, error) {
	hashedPassword, err := utils.Hash(c.Password)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
