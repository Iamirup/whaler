package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/application/ports"
	appService "github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/application/services"
	domainService "github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/domain/services"
	"github.com/Iamirup/whaler/backend/microservices/magazine/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
	Token  token.Token

	managmentApp *fiber.App
	clientApp    *fiber.App
}

func New(log *zap.Logger, newsRepo ports.NewsPersistencePort, token token.Token) *Server {
	server := &Server{Logger: log, Token: token}

	server.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	server.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	kubernetesHandler := NewKubernetesHandler(server)

	server.managmentApp.Get("/healthz/liveness", kubernetesHandler.Liveness)
	server.managmentApp.Get("/healthz/readiness", kubernetesHandler.Readiness)

	newsService := domainService.NewNewsService(newsRepo, log, token)
	newsHandler := NewNewsHandler(server, appService.NewNewsApplicationService(newsService, log))

	magazineV1 := server.clientApp.Group("/api/magazine/v1", newsHandler.fetchUserDataMiddleware)

	magazineV1.Post("/news", newsHandler.AddNews) // only for admins
	magazineV1.Get("/news", newsHandler.SeeNews)  // for both users and admins

	// 404 Handler
	magazineV1.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})

	return server
}

func (handler *Server) Serve() {
	go func() {
		err := handler.managmentApp.Listen(":8081")
		handler.Logger.Fatal("error resolving managment server", zap.Error(err))
	}()

	go func() {
		err := handler.clientApp.Listen(":8080")
		handler.Logger.Fatal("error resolving client server", zap.Error(err))
	}()
}
