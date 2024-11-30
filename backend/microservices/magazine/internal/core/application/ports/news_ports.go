package ports

import (
	serr "github.com/Iamirup/whaler/backend/microservices/magazine/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/domain/entity"
)

type (
	// NewsPersistencePort defines the methods for interacting with news data
	NewsPersistencePort interface {
		// CreateNews adds a new news to the database
		CreateNews(ticket *entity.News) error

		// GetNews retrieves all news from the database
		GetNews(encryptedCursor string, limit int) ([]entity.News, string, error)
	}

	// NewsServicePort defines the methods for interacting with news services
	NewsServicePort interface {
		AddNews(title, content string) (entity.UUID, *serr.ServiceError)
		SeeNews(encryptedCursor string, limit int) ([]entity.News, string, *serr.ServiceError)
	}
)
