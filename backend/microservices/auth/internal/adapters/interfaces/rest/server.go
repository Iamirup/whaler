package http

import (
	"encoding/json"

	"github.com/Iamirup/whaler/backend/auth/internal/repository"
	"github.com/Iamirup/whaler/backend/auth/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	logger     *zap.Logger
	repository repository.Repository
	token      token.Token

	managmentApp *fiber.App
	clientApp    *fiber.App
}

func New(log *zap.Logger, repo repository.Repository, token token.Token) *Server {
	server := &Server{logger: log, repository: repo, token: token}

	// Management Endpoints
	server.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	server.managmentApp.Get("/healthz/liveness", server.liveness)
	server.managmentApp.Get("/healthz/readiness", server.readiness)

	// Client Endpoints
	server.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	authV1 := server.clientApp.Group("/api/auth/v1")

	authV1.Post("/register", server.register)
	authV1.Post("/login", server.login)
	authV1.Post("/refresh", server.refresh)
	authV1.Post("/revoke", server.revoke)

	return server
}

func (handler *Server) Serve() {
	go func() {
		err := handler.managmentApp.Listen(":8081")
		handler.logger.Fatal("error resolving managment server", zap.Error(err))
	}()

	go func() {
		err := handler.clientApp.Listen(":8080")
		handler.logger.Fatal("error resolving client server", zap.Error(err))
	}()
}
