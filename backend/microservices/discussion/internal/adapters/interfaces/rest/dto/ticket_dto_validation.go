package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (v *NewCommentRequest) Validate() error {
	return validate.Struct(v)
}

func (v *GetCommentsRequest) Validate() error {
	return validate.Struct(v)
}
