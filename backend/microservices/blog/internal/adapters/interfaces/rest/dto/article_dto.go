package dto

import "github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"

// requests
type (
	NewArticleRequest struct {
		Title   string `json:"title"        form:"title"         validate:"required,min=3,max=70"`
		Content string `json:"content"      form:"content"       validate:"required,min=50,max=600"`
	}

	UpdateArticleRequest struct {
		ArticleId entity.UUID `json:"article_id"     form:"article_id"     validate:"required"`
		Title     string      `json:"title"          form:"title"          validate:"omitempty,min=3,max=70"`
		Content   string      `json:"content"        form:"content"        validate:"omitempty,min=50,max=600"`
	}

	DeleteArticleRequest struct {
		ArticleId entity.UUID `json:"article_id"     form:"article_id"        validate:"required"`
	}
)

// responses in successful status
type (
	NewArticleResponse struct {
		ArticleId entity.UUID `json:"article_id"  form:"article_id"`
	}

	UpdateArticleResponse struct {
		// nothing
	}

	DeleteArticleResponse struct {
		// nothing
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
