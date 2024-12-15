package repository

import (
	"errors"
	"fmt"
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

	hashedToken := entity.HashToken(refreshToken.Token, r.config.Pepper)

	in := []any{hashedToken, refreshToken.OwnerId}
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

const QueryCheckRefreshTokenInDBById = `
SELECT refresh_token
FROM refresh_tokens
WHERE owner_id = $1;`

func (r *refreshTokenRepository) CheckRefreshTokenInDBById(ownerId string) error {

	var refreshToken string

	in := []interface{}{ownerId}
	out := []interface{}{&refreshToken}
	if err := r.rdbms.QueryRow(QueryCheckRefreshTokenInDBById, in, out); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			r.logger.Error("Error id not found", zap.Error(err))
			return err
		}

		r.logger.Error("Error find refresh token by id", zap.Error(err))
		return err
	}

	return nil
}

const QueryRemoveRefreshToken = `
DELETE FROM refresh_tokens
WHERE refresh_token = $1;`

func (r *refreshTokenRepository) RemoveRefreshToken(refreshToken string) error {

	fmt.Println("refreshToken: ", refreshToken)

	hashedToken := entity.HashToken(refreshToken, r.config.Pepper)

	fmt.Println("hashedToken: ", hashedToken)

	in := []interface{}{hashedToken}
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

const QueryRevokeAllRefreshTokensById = `
DELETE FROM refresh_tokens
WHERE owner_id = $1;`

func (r *refreshTokenRepository) RevokeAllRefreshTokensById(userId string) error {

	in := []interface{}{userId}
	if err := r.rdbms.Execute(QueryRevokeAllRefreshTokensById, in); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			r.logger.Error("Error owner id not found", zap.Error(err))
			return err
		}

		r.logger.Error("Error find owner by its id", zap.Error(err))
		return err
	}

	return nil
}

const QueryCheckRefreshTokenExistsInDB = `
SELECT owner_id
FROM refresh_tokens
WHERE refresh_token = $1;`

func (r *refreshTokenRepository) CheckRefreshTokenExistsInDB(possibleRefreshToken string) string {

	var userId string

	in := []interface{}{possibleRefreshToken}
	out := []interface{}{&userId}
	if err := r.rdbms.QueryRow(QueryCheckRefreshTokenExistsInDB, in, out); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			r.logger.Error("Error id not found", zap.Error(err))
			return ""
		}

		r.logger.Error("Error find refresh token by id", zap.Error(err))
		return ""
	}

	return userId
}

const QueryCheckIfIsAdmin = `
SELECT COUNT(1)
FROM admins
WHERE id = $1;`

func (r *refreshTokenRepository) CheckIfIsAdmin(userId string) (bool, error) {

	var isAdmin bool

	in := []interface{}{userId}
	out := []interface{}{&isAdmin}
	if err := r.rdbms.QueryRow(QueryCheckIfIsAdmin, in, out); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			r.logger.Error("Error id not found", zap.Error(err))
			return false, err
		}

		r.logger.Error("Error find refresh token by id", zap.Error(err))
		return false, err
	}

	return isAdmin, nil
}
