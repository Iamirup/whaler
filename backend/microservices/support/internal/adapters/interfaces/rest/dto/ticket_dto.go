package dto

import "github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"

// requests
type (
	NewTicketRequest struct {
		Title   string `json:"title"        form:"title"         validate:"required,min=3,max=80"`
		Content string `json:"content"      form:"content"       validate:"required,min=60,max=600"`
	}

	MyTicketsRequest struct {
		// recognizes from jwt data (user id)
	}

	ReplyToTicketRequest struct {
		TicketId  entity.uuid `json:"ticket_id"     form:"ticket_id"        validate:"required"`
		ReplyText string      `json:"reply_text"    form:"reply_text"       validate:"required,min=40,max=500"`
	}

	AllTicketRequest struct {
		// recognizes from jwt data (user id)
	}
)

// responses in successful status
type (
	NewTicketResponse struct {
		TicketId entity.uuid `json:"ticket_id"  form:"ticket_id"`
	}

	MyTicketsResponse struct {
		Tikcets []entity.Ticket `json:"tickets"   form:"tickets"`
	}

	ReplyToTicketResponse struct {
		// nothing
	}

	AllTicketResponse struct {
		Tikcets []entity.Ticket `json:"tickets"   form:"tickets"`
	}
)

// responses in unsuccessful status
type (
	ErrorResponse struct {
		Error       string `json:"error"         form:"error"`
		NeedRefresh bool   `json:"need_refresh"  form:"need_refresh"`
	}
)
