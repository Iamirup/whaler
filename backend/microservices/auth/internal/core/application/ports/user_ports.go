package ports

import (
	"context"
	"database/sql"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
)

type (
	// UserPersistencePort defines the methods for interacting with user data
	UserPersistencePort interface {
		// CreateUser adds a new user to the database
		CreateUser(user *entity.User) error

		// GetUserByUsername retrieves a user by its username
		GetUserByUsername(username string) (*entity.User, error)

		// GetUserByUsernameAndPassword retrieves a user by its username and password
		GetUserByUsernameAndPassword(username, password string) (*entity.User, error)
	}

	// UserServicePort defines the methods for interacting with user services
	UserServicePort interface {
		Register(ctx context.Context, code string, maxUsages int) (*serr.ServiceError, entity.AuthTokens)
		Login(ctx context.Context, code string, tx *sql.Tx) (*serr.ServiceError, entity.AuthTokens)
	}
)
