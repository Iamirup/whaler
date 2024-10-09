package http

import (
	"encoding/json"

	"github.com/Iamirup/whaler/internal/repository"
	"github.com/Iamirup/whaler/pkg/token"
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

	apiV1 := server.clientApp.Group("/api/v1")

	auth := apiV1.Group("/auth")
	auth.Post("/register", server.register)
	auth.Post("/login", server.login)

	events := apiV1.Group("/events", server.fetchUserId)
	events.Get("", server.events)

	configs := apiV1.Group("/config", server.fetchUserId)
	configs.Get("", server.getConfig)
	configs.Patch("", server.updateConfig)

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
