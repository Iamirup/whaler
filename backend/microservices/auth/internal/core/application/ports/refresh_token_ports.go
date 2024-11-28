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

		// CheckRefreshTokenInDBById checks a refresh token is valid in database by its owener user id
		CheckRefreshTokenInDBById(userId, refreshToken string) error

		// RemoveRefreshTokenById removes a refresh token by its owener user id from database
		RemoveRefreshToken(userId string) error

		// RevokeAllRefreshTokensById removes all refresh token which is related to its owener user id
		RevokeAllRefreshTokensById(userId string) error

		// CheckIfIsAdmin checks if a user is admin or not
		CheckIfIsAdmin(userId string) (bool, error)
	}

	// RefreshTokenServicePort defines the methods for interacting with refresh token services
	RefreshTokenServicePort interface {
		CheckRefreshTokenInDBById(userId, refreshToken string) *serr.ServiceError
		RevokeAllRefreshTokensById(userId string) error
		CheckIfIsAdmin(userId string) (bool, error)
	}
)
