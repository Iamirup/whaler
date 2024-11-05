package repository

import (
	"errors"
	"strings"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/rdbms"
	"go.uber.org/zap"
)

const QueryCreateNewRefreshToken = `
INSERT INTO refresh_tokens(refresh_token, owner_id)
VALUES($1, $2)
RETURNING owner_id;`

func (r *refreshTokenRepository) CreateNewRefreshToken(refreshToken *entity.RefreshToken) error {

	if len(refreshToken.Token) == 0 {
		return errors.New("insufficient refresh token")
	}

	in := []any{refreshToken.Token, refreshToken.OwnerId}
	out := []any{&refreshToken.OwnerId}
	if err := r.rdbms.QueryRow(QueryCreateNewRefreshToken, in, out); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			r.logger.Error(err.Error(), zap.Error(err))
			return err
		}
		r.logger.Error("Error inserting new refresh token", zap.Error(err))
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
			r.logger.Error("Error id not found", zap.Error(err))
			return nil, err
		}

		r.logger.Error("Error find refresh token by id", zap.Error(err))
		return nil, err
	}

	return refreshToken, nil
}

const QueryRemoveRefreshToken = `
DELETE FROM refresh_tokens
WHERE refresh_token = $1;`

func (r *refreshTokenRepository) RemoveRefreshToken(refreshToken string) error {

	in := []interface{}{refreshToken}
	if err := r.rdbms.Execute(QueryRemoveRefreshToken, in); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			r.logger.Error("Error owner refresh token not found", zap.Error(err))
			return err
		}

		r.logger.Error("Error find refresh token by owner id", zap.Error(err))
		return err
	}

	return nil
}
