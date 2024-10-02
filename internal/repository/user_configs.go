package repository

import (
	"github.com/Iamirup/whaler/internal/models"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

const QueryGetUserConfigById = `
SELECT *
FROM user_configs
WHERE user_id = $1 AND id = $2;`

func (r *repository) GetConfigById(userId, userConfigId uint64) (*models.UserConfig, error) {

	userConfig := models.UserConfig{Id: userConfigId}

	in := []any{userId, userConfigId}
	out := []any{&userConfig.Name, pq.Array(&userConfig.Phones), &userConfig.Description}
	if err := r.rdbms.QueryRow(QueryGetUserConfigById, in, out); err != nil {
		r.logger.Error("Error get userConfig by id", zap.Error(err))
		return nil, err
	}

	return &userConfig, nil
}

const QueryUpdateUserConfig = `
UPDATE user_configs
SET name = $1, phones = $2, description = $3
WHERE user_id = $4 AND id = $5;`

func (r *repository) UpdateConfig(userId uint64, userConfig *models.UserConfig) error {

	in := []any{userConfig.Name, pq.Array(userConfig.Phones), userConfig.Description, userId, userConfig.Id}
	if err := r.rdbms.Execute(QueryUpdateUserConfig, in); err != nil {
		r.logger.Error("Error updating userConfig", zap.Error(err))
		return err
	}

	return nil
}
