package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (v *AddNewsRequest) Validate() error {
	return validate.Struct(v)
}

func (v *SeeNewsRequest) Validate() error {
	return validate.Struct(v)
}
