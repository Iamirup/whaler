package dto

import "github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/domain/entity"

// requests
type (
	AddNewsRequest struct {
		Title   string `json:"title"        form:"title"         validate:"required,min=5,max=100"`
		Content string `json:"content"      form:"content"       validate:"required,min=100,max=2000"`
	}

	SeeNewsRequest struct {
		// nothing (just cursor and limit in query parameters)
	}
)

// responses in successful status
type (
	AddNewsResponse struct {
		NewsId entity.UUID `json:"news_id"  form:"news_id"`
	}

	SeeNewsResponse struct {
		News      []entity.News `json:"news"         form:"news"`
		NewCursor string        `json:"new_cursor"   form:"new_cursor"`
	}
)

// responses in unsuccessful status
type (
	ErrorContent struct {
		Field   string `json:"field"     form:"field"`
		Message string `json:"message"   form:"message"`
	}

	ErrorResponse struct {
		Errors      []ErrorContent `json:"errors"         form:"errors"`
		NeedRefresh bool           `json:"need_refresh"   form:"need_refresh"`
	}
)
