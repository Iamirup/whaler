package dto

import "github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"

// requests
type (
	NewTicketRequest struct {
		Title   string `json:"title"        form:"title"         validate:"required,min=3,max=80"`
		Content string `json:"content"      form:"content"       validate:"required,min=60,max=600"`
	}

	MyTicketsRequest struct {
		Cursor string `json:"cursor"     form:"cursor"      validate:"required"`
		Limit  int    `json:"limit"      form:"limit"       validate:"required"`
	}

	ReplyToTicketRequest struct {
		TicketId  entity.UUID `json:"ticket_id"     form:"ticket_id"        validate:"required"`
		ReplyText string      `json:"reply_text"    form:"reply_text"       validate:"required,min=40,max=500"`
	}

	AllTicketRequest struct {
		Cursor string `json:"cursor"     form:"cursor"      validate:"required"`
		Limit  int    `json:"limit"      form:"limit"       validate:"required"`
	}
)

// responses in successful status
type (
	NewTicketResponse struct {
		TicketId entity.UUID `json:"ticket_id"  form:"ticket_id"`
	}

	MyTicketsResponse struct {
		Tickets   []entity.Ticket `json:"tickets"   form:"tickets"`
		NewCursor string          `json:"new_cursor"   form:"new_cursor"`
	}

	ReplyToTicketResponse struct {
		// nothing
	}

	AllTicketResponse struct {
		Tickets   []entity.Ticket `json:"tickets"   form:"tickets"`
		NewCursor string          `json:"new_cursor"   form:"new_cursor"`
	}
)

// responses in unsuccessful status
type (
	ErrorResponse struct {
		Error       string `json:"error"         form:"error"`
		NeedRefresh bool   `json:"need_refresh"  form:"need_refresh"`
	}
)
