package ports

import (
	serr "github.com/Iamirup/whaler/backend/microservices/blog/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"
)

type (
	// ArticlePersistencePort defines the methods for interacting with article data
	ArticlePersistencePort interface {
		// CreateArticle adds a new article to the database
		CreateArticle(article *entity.Article) error

		// GetAnArticle retrieves the article that matches the url path
		GetAnArticle(articleId entity.UUID) (*entity.Article, error)

		// GetAllArticles retrieves all articles in the system
		GetAllArticles(encryptedCursor string, limit int) ([]entity.Article, string, error)

		// GetMyArticles retrieves all articles in the system that belongs to this author
		GetMyArticles(encryptedCursor string, limit int, authorId entity.UUID) ([]entity.Article, string, error)

		// UpdateArticleTitle updates article title in database
		UpdateArticleTitle(articleId entity.UUID, title string) error

		// UpdateArticleContent updates article content in database
		UpdateArticleContent(articleId entity.UUID, content string) error

		// RemoveArticle deletes a specific article from database
		RemoveArticle(articleId entity.UUID) error

		// CheckIfIsAuthorById checks if the the user is the same author of the article or not
		CheckIfIsAuthorById(articleId, authorId entity.UUID) error
	}

	// ArticleServicePort defines the methods for interacting with article services
	ArticleServicePort interface {
		GetAnArticle(articleId entity.UUID) (*entity.Article, *serr.ServiceError)
		GetAllArticles(encryptedCursor string, limit int) ([]entity.Article, string, *serr.ServiceError)
		GetMyArticles(encryptedCursor string, limit int, authorId entity.UUID) ([]entity.Article, string, *serr.ServiceError)
		NewArticle(title, content string, authorId entity.UUID, authorUsername string) (entity.UUID, *serr.ServiceError)
		UpdateArticle(articleId entity.UUID, title, content string, authorId entity.UUID) *serr.ServiceError
		DeleteArticle(articleId, authorId entity.UUID) *serr.ServiceError
	}
)
