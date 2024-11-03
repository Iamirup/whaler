package repository

import (
	"errors"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/rdbms"
	"go.uber.org/zap"
)

const QueryCreateNewRefreshToken = `
INSERT INTO refresh_tokens(refresh_token, owner_id)
VALUES($1, $2)
ON CONFLICT (owner_id) 
DO UPDATE SET refresh_token = EXCLUDED.refresh_token
RETURNING owner_id;`

func (r *refreshTokenRepository) CreateNewRefreshToken(refreshToken *entity.RefreshToken) error {

	if len(refreshToken.Token) == 0 {
		return errors.New("insufficient refresh token")
	}

	in := []any{refreshToken.Token, refreshToken.OwnerId}
	out := []any{&refreshToken.OwnerId}
	if err := r.rdbms.QueryRow(QueryCreateNewRefreshToken, in, out); err != nil {
		r.logger.Error("Error inserting author", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetRefreshTokenById = `
SELECT refresh_token
FROM refresh_tokens
WHERE owner_id = $1;`

func (r *refreshTokenRepository) GetRefreshTokenById(ownerId string) (*entity.RefreshToken, error) {

	refreshToken := &entity.RefreshToken{OwnerId: ownerId}

	in := []interface{}{ownerId}
	out := []interface{}{&refreshToken.Token}
	if err := r.rdbms.QueryRow(QueryGetRefreshTokenById, in, out); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			return nil, err
		}

		r.logger.Error("Error find user by username", zap.Error(err))
		return nil, err
	}

	return refreshToken, nil
}

const QueryRemoveRefreshTokenById = `
DELETE FROM refresh_tokens
WHERE refresh_token = $1
RETURNING refresh_token;`

func (r *refreshTokenRepository) RemoveRefreshTokenById(refreshToken string) error {

	refreshTokenEntity := &entity.RefreshToken{Token: refreshToken}

	in := []interface{}{refreshToken}
	out := []interface{}{&refreshTokenEntity.Token}
	if err := r.rdbms.QueryRow(QueryRemoveRefreshTokenById, in, out); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			return err
		}

		r.logger.Error("Error find refresh token by owner id", zap.Error(err))
		return err
	}

	return nil
}
