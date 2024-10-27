package ports

import (
	"context"
	"database/sql"
)

type (
	// RefreshTokenPersistencePort defines the methods for interacting with refresh token data
	RefreshTokenPersistencePort interface {
		// CreateNewRefreshToken adds a new refresh token to the database
		CreateNewRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error

		// GetRefreshTokenById retrieves a refresh token by its owener user id from database
		GetRefreshTokenById(ctx context.Context, id uuid) (*entity.RefreshToken, error)
	}

	// RefreshTokenServicePort defines the methods for interacting with refresh token services
	RefreshTokenServicePort interface {
		Refresh(ctx context.Context, code string, maxUsages int) error
		Revoke(ctx context.Context, code string, tx *sql.Tx) (*entity.VoucherCode, error)
	}
)
