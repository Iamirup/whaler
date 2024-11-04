package dto

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func usernameValidator(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{2,31}$`)
	return re.MatchString(fl.Field().String())
}

func passwordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 255 {
		return false
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password) {
		return false
	}
	return true
}

func init() {
	validate.RegisterValidation("username", usernameValidator)
	validate.RegisterValidation("password", passwordValidator)
}

func (v *LoginRequest) Validate() error {
	return validate.Struct(v)
}

func (v *RegisterRequest) Validate() error {
	return validate.Struct(v)
}
