package repository

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/support/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/support/pkg/utils"
	"go.uber.org/zap"
)

type ticketRepository struct {
	logger *zap.Logger
	config *Config
	rdbms  rdbms.RDBMS
}

type MigrationRepository interface {
	Migrate(entity.Migrate) error
}

type migrationRepository struct {
	logger *zap.Logger
	config *Config
	rdbms  rdbms.RDBMS
}

func NewTicketRepository(logger *zap.Logger, cfg *Config, rdbms rdbms.RDBMS) ports.TicketPersistencePort {
	return &ticketRepository{logger: logger, config: cfg, rdbms: rdbms}
}

func NewMigrationRepository(logger *zap.Logger, cfg *Config, rdbms rdbms.RDBMS) MigrationRepository {
	return &migrationRepository{logger: logger, config: cfg, rdbms: rdbms}
}

//go:embed migrations
var migrations embed.FS

func (r *migrationRepository) Migrate(direction entity.Migrate) error {

	files, err := fs.ReadDir(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("error reading migrations directory:\n%v", err)
	}

	result := make([]string, 0, len(files)/2)

	for _, file := range files {
		splits := strings.Split(file.Name(), ".")
		if splits[1] == string(direction) {
			result = append(result, file.Name())
		}
	}

	result = utils.Sort(result)

	for index := 0; index < len(result); index++ {
		file := "migrations/"

		if direction == entity.Up {
			file += result[index]
		} else {
			file += result[len(result)-index-1]
		}

		data, err := fs.ReadFile(migrations, file)
		if err != nil {
			return fmt.Errorf("error reading migration file: %s\n%v", file, err)
		}

		if err := r.rdbms.Execute(string(data), []any{}); err != nil {
			return fmt.Errorf("error migrating the file: %s\n%v", file, err)
		}
	}

	return nil
}
