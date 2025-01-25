package cmd

import (
	"os"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/adapters/interfaces/rest"
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/config"
	"github.com/Iamirup/whaler/backend/microservices/eventor/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/eventor/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/eventor/pkg/token"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run eventor server",
		Run:   run,
	}
}

func (cmd *Server) main(cfg *config.Config, trap chan os.Signal) {
	myLogger := logger.NewZap(cfg.Logger)

	db, err := rdbms.New(cfg.RDBMS)
	if err != nil {
		myLogger.Panic("Error creating rdbms database", zap.Error(err))
	}

	tableConfigRepo := repository.NewTableConfigRepository(myLogger, cfg.Repository, db)

	theToken, err := token.New(cfg.Token)
	if err != nil {
		myLogger.Panic("Error creating token object", zap.Error(err))
	}

	rest.New(myLogger, tableConfigRepo, theToken).Serve()

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	myLogger.Info("exiting by receiving a unix signal", field)
}
