package ports

import (
	serr "github.com/Iamirup/whaler/backend/microservices/discussion/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/domain/entity"
)

type (
	// CommentPersistencePort defines the methods for interacting with comment data
	CommentPersistencePort interface {
		// AddComment adds a new comment to the database
		AddComment(comment *entity.Comment) error

		// GetComments retrieves all comments
		GetComments(currencyTopic, encryptedCursor string, limit int) ([]entity.Comment, string, error)
	}

	// CommentServicePort defines the methods for interacting with comment services
	CommentServicePort interface {
		NewComment(text, currency string, username string) (int64, *serr.ServiceError)
		GetComments(currencyTopic, encryptedCursor string, limit int) ([]entity.Comment, string, *serr.ServiceError)
	}
)
