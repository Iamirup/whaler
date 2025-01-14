package dto

import "github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/domain/entity"

// requests
type (
	NewCommentRequest struct {
		Currency string `json:"currency"    form:"currency"     validate:"required"`
		Text     string `json:"text"        form:"text"         validate:"required,max=600"`
	}

	GetCommentsRequest struct {
		// nothing (just cursor and limit in query parameters)
	}
)

// responses in successful status
type (
	NewCommentResponse struct {
		CommentId entity.UUID `json:"comment_id"  form:"comment_id"`
	}

	GetCommentsResponse struct {
		Comments    []entity.Comment `json:"comments"      form:"comments"`
		NewCursor   string           `json:"new_cursor"    form:"new_cursor"`
		OwnUsername string           `json:"own_username"    form:"own_username"     validate:"required"`
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
