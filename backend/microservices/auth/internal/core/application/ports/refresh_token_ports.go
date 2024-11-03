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

		// GetRefreshTokenById retrieves a refresh token by its owener user id from database
		GetRefreshTokenById(userId string) (*entity.RefreshToken, error)

		// RemoveRefreshTokenById removes a refresh token by its owener user id from database
		RemoveRefreshTokenById(userId string) error
	}

	// RefreshTokenServicePort defines the methods for interacting with refresh token services
	RefreshTokenServicePort interface {
		Refresh(refreshToken string) *serr.ServiceError
		GetRefreshTokenById(userId string) (*entity.RefreshToken, *serr.ServiceError)
	}
)
