package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (v *UpdateTableConfigRequest) Validate() error {
	return validate.Struct(v)
}

func (v *SeeTableRequest) Validate() error {
	return validate.Struct(v)
}
