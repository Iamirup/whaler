package dto

import "github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"

// requests
type (
	NewTicketRequest struct {
		Title   string `json:"title"        form:"title"         validate:"required,min=3,max=70"`
		Content string `json:"content"      form:"content"       validate:"required,min=50,max=600"`
	}

	MyTicketsRequest struct {
		// nothing (just cursor and limit in query parameters)
	}

	ReplyToTicketRequest struct {
		TicketId  entity.UUID `json:"ticket_id"     form:"ticket_id"        validate:"required"`
		ReplyText string      `json:"reply_text"    form:"reply_text"       validate:"required,min=30,max=800"`
	}

	AllTicketRequest struct {
		// nothing (just cursor and limit in query parameters)
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
	ErrorContent struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ErrorResponse struct {
		Errors    []ErrorContent `json:"errors"`
		NeedLogin bool           `json:"need_login"`
	}
)
