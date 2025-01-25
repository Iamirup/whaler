package ports

import (
	serr "github.com/Iamirup/whaler/backend/microservices/eventor/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/domain/entity"
)

type (
	// TableConfigPersistencePort defines the methods for interacting with tableConfig data
	TableConfigPersistencePort interface {
		// UpdateTableConfig updates previous table config to the database
		UpdateTableConfig(ticket *entity.TableConfig) error

		// GetTableConfig retrieves the user table config from the database
		GetTableConfig(encryptedCursor string, limit int) ([]entity.TableConfig, string, error)
	}

	// TableConfigServicePort defines the methods for interacting with tableConfig services
	TableConfigServicePort interface {
		UpdateTableConfig(title, content string) (entity.TableConfig, *serr.ServiceError)
	}
)
