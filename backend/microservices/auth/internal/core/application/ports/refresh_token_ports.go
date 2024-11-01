package ports

import (
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	"github.com/gofiber/fiber/v2"
)

type (
	// RefreshTokenPersistencePort defines the methods for interacting with refresh token data
	RefreshTokenPersistencePort interface {
		// CreateNewRefreshToken adds a new refresh token to the database
		CreateNewRefreshToken(refreshToken *entity.RefreshToken) error

		// GetRefreshTokenById retrieves a refresh token by its owener user id from database
		GetRefreshTokenById(id string) (*entity.RefreshToken, error)
	}

	// RefreshTokenServicePort defines the methods for interacting with refresh token services
	RefreshTokenServicePort interface {
		Refresh(ctx *fiber.Ctx, something string) error
		Revoke(ctx *fiber.Ctx, something string) error
	}
)
