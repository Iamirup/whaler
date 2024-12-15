package ports

import (
	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
)

type (
	// UserPersistencePort defines the methods for interacting with user data
	UserPersistencePort interface {
		// CreateUser adds a new user to the database
		CreateUser(user *entity.User) error

		// GetUserByUsernameAndPassword retrieves a user by its username and password
		GetUserByUsernameAndPassword(username, password string) (*entity.User, error)

		// GetUserByEmailAndPassword retrieves a user by its email and password
		GetUserByEmailAndPassword(email, password string) (*entity.User, error)
	}

	// UserServicePort defines the methods for interacting with user services
	UserServicePort interface {
		Register(email, username, password string) (entity.AuthTokens, *serr.ServiceError)
		Login(email, username, password, possibleRefreshToken string) (entity.AuthTokens, *serr.ServiceError)
		Logout(refreshToken string) *serr.ServiceError
	}
)
