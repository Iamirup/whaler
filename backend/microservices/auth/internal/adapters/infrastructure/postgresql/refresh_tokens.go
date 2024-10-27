package repository

import (
	"errors"

	"github.com/Iamirup/whaler/backend/eventor/internal/models"
	"github.com/Iamirup/whaler/backend/eventor/pkg/rdbms"
	"go.uber.org/zap"
)

const QueryCreateNewRefreshToken = `
INSERT INTO refresh_tokens(refresh_token, owner_id)
VALUES($1, $2)
ON CONFLICT (owner_id) 
DO UPDATE SET refresh_token = EXCLUDED.refresh_token
RETURNING owner_id;`

func (r *repository) CreateNewRefreshToken(refreshToken *models.RefreshToken) error {

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

func (r *repository) GetRefreshTokenById(ownerId string) (*models.RefreshToken, error) {

	refreshToken := &models.RefreshToken{OwnerId: ownerId}

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
