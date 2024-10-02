package cmd

import (
	"os"

	"github.com/mohammadne/phone-book/internal/api/http"
	"github.com/mohammadne/phone-book/internal/config"
	"github.com/mohammadne/phone-book/internal/repository"
	"github.com/mohammadne/phone-book/pkg/logger"
	"github.com/mohammadne/phone-book/pkg/rdbms"
	"github.com/mohammadne/phone-book/pkg/token"
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
		Short: "run floki server",
		Run:   run,
	}
}

func (cmd *Server) main(cfg *config.Config, trap chan os.Signal) {
	myLogger := logger.NewZap(cfg.Logger)

	db, err := rdbms.New(cfg.RDBMS)
	if err != nil {
		myLogger.Panic("Error creating rdbms database", zap.Error(err))
	}

	repo := repository.New(myLogger, cfg.Repository, db)

	theToken, err := token.New(cfg.Token)
	if err != nil {
		myLogger.Panic("Error creating token object", zap.Error(err))
	}

	http.New(myLogger, repo, theToken).Serve()

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	myLogger.Info("exiting by receiving a unix signal", field)
}
