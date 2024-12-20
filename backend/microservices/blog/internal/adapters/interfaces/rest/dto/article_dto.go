package dto

import "github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"

// requests
type (
	GetAnArticleRequest struct {
		// nothing (just cursor and limit in query parameters)
	}

	GetArticlesRequest struct {
		// nothing (just cursor and limit in query parameters)
	}

	NewArticleRequest struct {
		Title   string `json:"title"        form:"title"         validate:"required,min=3,max=70"`
		Content string `json:"content"      form:"content"       validate:"required,min=50"`
	}

	UpdateArticleRequest struct {
		ArticleId entity.UUID `json:"article_id"     form:"article_id"     validate:"required"`
		Title     string      `json:"title"          form:"title"          validate:"omitempty,min=3,max=70"`
		Content   string      `json:"content"        form:"content"        validate:"omitempty,min=50"`
	}

	DeleteArticleRequest struct {
		ArticleId entity.UUID `json:"article_id"     form:"article_id"        validate:"required"`
	}

	LikeArticleRequest struct {
		ArticleId entity.UUID `json:"article_id"     form:"article_id"     validate:"required"`
	}

	GetTopAuthorsRequest struct {
		// nothing
	}

	GetPopularArticlesRequest struct {
		// nothing
	}
)

// responses in successful status
type (
	GetAnArticleResponse struct {
		Article entity.Article `json:"article"     form:"article"`
	}

	GetArticlesResponse struct {
		Articles  []entity.Article `json:"articles"     form:"articles"`
		NewCursor string           `json:"new_cursor"   form:"new_cursor"`
	}

	NewArticleResponse struct {
		ArticleId entity.UUID `json:"article_id"  form:"article_id"`
	}

	UpdateArticleResponse struct {
		// nothing
	}

	DeleteArticleResponse struct {
		// nothing
	}

	LikeArticleResponse struct {
		// nothing
	}

	GetTopAuthorsResponse struct {
		TopAuthors []entity.TopAuthor `json:"authors"  form:"authors"`
	}

	GetPopularArticlesResponse struct {
		Articles []entity.Article `json:"articles"     form:"articles"`
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
