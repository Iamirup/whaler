package dto

import (
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

func (v *LoginRequest) Validate() error {
	return validate.Struct(v)
}

func (v *RegisterRequest) Validate() error {
	return validate.Struct(v)
}
