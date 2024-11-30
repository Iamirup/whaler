package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (v *NewTicketRequest) Validate() error {
	return validate.Struct(v)
}

func (v *MyTicketsRequest) Validate() error {
	return validate.Struct(v)
}

func (v *ReplyToTicketRequest) Validate() error {
	return validate.Struct(v)
}

func (v *AllTicketRequest) Validate() error {
	return validate.Struct(v)
}
