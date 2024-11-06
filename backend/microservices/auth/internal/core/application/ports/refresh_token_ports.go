package ports

import (
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
)

type (
	// RefreshTokenPersistencePort defines the methods for interacting with refresh token data
	RefreshTokenPersistencePort interface {
		// CreateNewRefreshToken adds a new refresh token to the database
		CreateNewRefreshToken(refreshToken *entity.RefreshToken) error

		// GetAndCheckRefreshTokenById retrieves a refresh token by its owener user id from database
		GetAndCheckRefreshTokenById(userId, refreshToken string) error

		// RemoveRefreshTokenById removes a refresh token by its owener user id from database
		RemoveRefreshToken(userId string) error
	}

	// RefreshTokenServicePort defines the methods for interacting with refresh token services
	RefreshTokenServicePort interface {
		GetAndCheckRefreshTokenById(userId, refreshToken string) *serr.ServiceError
	}
)
