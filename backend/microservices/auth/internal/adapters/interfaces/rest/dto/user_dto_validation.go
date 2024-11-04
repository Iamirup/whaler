package dto

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func usernameValidator(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	if len(str) < 3 || len(str) > 32 {
		return false
	}
	if regexp.MustCompile(`^[_.]`).MatchString(str) || regexp.MustCompile(`[_.]$`).MatchString(str) {
		return false
	}
	if regexp.MustCompile(`[_.]{2}`).MatchString(str) {
		return false
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9._]+$`).MatchString(str) {
		return false
	}
	return true
}

func strongPasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 255 {
		return false
	}
	if !regexp.MustCompile(`[A-Za-z]`).MatchString(password) {
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
	validate.RegisterValidation("strong_password", strongPasswordValidator)
}

func (v *LoginRequest) Validate() error {
	return validate.Struct(v)
}

func (v *RegisterRequest) Validate() error {
	return validate.Struct(v)
}
