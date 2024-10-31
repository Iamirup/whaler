package rest

import (
	"encoding/json"

	"github.com/Iamirup/whaler/backend/microservice/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/handlers"
	appService "github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/services"
	domainService "github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/services"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"
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

func New(log *zap.Logger, userRepo ports.UserPersistencePort, refreshTokenRepo ports.RefreshTokenPersistencePort, token token.Token) *Server {
	server := &Server{logger: log, repository: repo, token: token}

	// Management Endpoints
	server.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	kubernetesHandler := handlers.KubernetesHandler{server: server}

	server.managmentApp.Get("/healthz/liveness", kubernetesHandler.liveness)
	server.managmentApp.Get("/healthz/readiness", kubernetesHandler.readiness)

	// Client Endpoints
	server.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	authV1 := server.clientApp.Group("/api/auth/v1")

	userService := domainService.NewUserService(userRepo, refreshTokenRepo, log, token)
	userHandler := handlers.UserHandler{server: server, userAppService: appService.NewUserApplicationService(userService)}

	refreshTokenService := domainService.NewRefreshTokenService(refreshTokenRepo, log, token)
	refreshTokenHandler := handlers.RefreshTokenHandler{server: server, userAppService: appService.NewUserApplicationService(refreshTokenService)}

	authV1.Post("/register", userHandler.register)
	authV1.Post("/login", userHandler.login)
	authV1.Post("/refresh", refreshTokenHandler.refresh)
	authV1.Post("/revoke", refreshTokenHandler.revoke)

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
