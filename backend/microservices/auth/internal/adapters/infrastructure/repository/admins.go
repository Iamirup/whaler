package repository

import (
	"strings"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/rdbms"
	"go.uber.org/zap"
)

const QueryCreateNewAdmin = `
INSERT INTO admins(id) VALUES($1)
RETURNING id;`

func (r *adminRepository) CreateNewAdmin(userId entity.UUID) error {

	var adminId string

	in := []any{userId}
	out := []any{&adminId}
	if err := r.rdbms.QueryRow(QueryCreateNewAdmin, in, out); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			r.logger.Error(err.Error(), zap.Error(err))
			return err
		}
		r.logger.Error("Error inserting user", zap.Error(err))
		return err
	}

	return nil
}

const QueryRemoveAdmin = `
DELETE FROM admins
WHERE id = $1;`

func (r *adminRepository) RemoveAdmin(adminId entity.UUID) error {

	in := []interface{}{adminId}
	if err := r.rdbms.Execute(QueryRemoveAdmin, in); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			r.logger.Error("Error admin not found", zap.Error(err))
			return err
		}

		r.logger.Error("Error find admin by his id", zap.Error(err))
		return err
	}

	return nil
}
