package services

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/eventor/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/microservices/eventor/pkg/errors"
	"go.uber.org/zap"
)

type TableConfigService struct {
	tableConfigPersistencePort ports.TableConfigPersistencePort
	logger                     *zap.Logger
	token                      token.Token
}

func NewTableConfigService(
	tableConfigPersistencePort ports.TableConfigPersistencePort,
	logger *zap.Logger, token token.Token) ports.TableConfigServicePort {
	return &TableConfigService{
		tableConfigPersistencePort: tableConfigPersistencePort,
		logger:                     logger,
		token:                      token,
	}
}

func (s *TableConfigService) UpdateTableConfig(title, content string) (entity.TableConfig, *serr.ServiceError) {

	// tableConfigEntity := &entity.TableConfig{}

	// err := s.tableConfigPersistencePort.CreateTableConfig(tableConfigEntity)
	// if err != nil {
	// 	s.logger.Error("Error happened while creating the tableConfig", zap.Error(err))
	// 	return "", &serr.ServiceError{Message: "Error happened while creating the tableConfig", StatusCode: http.StatusInternalServerError}
	// } else if tableConfigEntity.TableConfigId == "" {
	// 	s.logger.Error("Error invalid tableConfig id created", zap.Any("tableConfig", tableConfigEntity))
	// 	return "", &serr.ServiceError{Message: "Error invalid tableConfig id created", StatusCode: http.StatusInternalServerError}
	// }

	// return tableConfigEntity.TableConfigId, nil
	return entity.TableConfig{}, nil
}

func (s *TableConfigService) SeeTable(encryptedCursor string, limit int) ([]entity.TableConfig, string, *serr.ServiceError) {
	tableConfig, newEncryptedCursor, err := s.tableConfigPersistencePort.GetTableConfig(encryptedCursor, limit)
	if err != nil {
		s.logger.Error("Something went wrong in retrieving tableConfig", zap.Error(err))
		return nil, "", &serr.ServiceError{Message: "Something went wrong in retrieving tableConfig", StatusCode: http.StatusInternalServerError}
	}

	return tableConfig, newEncryptedCursor, nil
}
