package cmd

import (
	"os"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/adapters/interfaces/rest"
	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/config"
	"github.com/Iamirup/whaler/backend/microservices/magazine/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/magazine/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/magazine/pkg/token"
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
		Short: "run magazine server",
		Run:   run,
	}
}

func (cmd *Server) main(cfg *config.Config, trap chan os.Signal) {
	myLogger := logger.NewZap(cfg.Logger)

	db, err := rdbms.New(cfg.RDBMS)
	if err != nil {
		myLogger.Panic("Error creating rdbms database", zap.Error(err))
	}

	newsRepo := repository.NewNewsRepository(myLogger, cfg.Repository, db)

	theToken, err := token.New(cfg.Token)
	if err != nil {
		myLogger.Panic("Error creating token object", zap.Error(err))
	}

	rest.New(myLogger, newsRepo, theToken).Serve()

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	myLogger.Info("exiting by receiving a unix signal", field)
}
