package cmd

import (
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/config"
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/rdbms"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Migrate struct{}

func (m Migrate) Command() *cobra.Command {
	run := func(_ *cobra.Command, args []string) {
		m.main(config.Load(true), args)
	}

	return &cobra.Command{
		Use:       "migrate",
		Short:     "run migrations",
		Run:       run,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"up", "down"},
	}
}

func (m *Migrate) main(cfg *config.Config, args []string) {
	myLogger := logger.NewZap(cfg.Logger)

	if len(args) != 1 {
		myLogger.Fatal("Invalid arguments given", zap.Any("args", args))
	}

	db, err := rdbms.New(cfg.RDBMS)
	if err != nil {
		myLogger.Fatal("Error creating rdbms", zap.Error(err))
	}

	repo := repository.NewMigrationRepository(myLogger, cfg.Repository, db)
	if err := repo.Migrate(entity.Migrate(args[0])); err != nil {
		myLogger.Fatal("Error migrating", zap.String("migration", args[0]), zap.Error(err))
	}

	myLogger.Info("Database has been migrated successfully", zap.String("migration", args[0]))
}
