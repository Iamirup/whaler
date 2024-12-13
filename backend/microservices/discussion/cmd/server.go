package cmd

import (
	"os"

	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/adapters/interfaces/rest"
	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/config"
	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/token"
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
		Short: "run discussion server",
		Run:   run,
	}
}

func (cmd *Server) main(cfg *config.Config, trap chan os.Signal) {
	myLogger := logger.NewZap(cfg.Logger)

	db, err := rdbms.New(cfg.RDBMS)
	if err != nil {
		myLogger.Panic("Error creating rdbms database", zap.Error(err))
	}

	ticketRepo := repository.NewTicketRepository(myLogger, cfg.Repository, db)

	theToken, err := token.New(cfg.Token)
	if err != nil {
		myLogger.Panic("Error creating token object", zap.Error(err))
	}

	rest.New(myLogger, ticketRepo, theToken).Serve()

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	myLogger.Info("exiting by receiving a unix signal", field)
}
