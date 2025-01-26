package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/ports"
	appService "github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/services"
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/services/transaction"
	domainService "github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/domain/services"
	"github.com/Iamirup/whaler/backend/microservices/eventor/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
	Token  token.Token

	managmentApp *fiber.App
	clientApp    *fiber.App
}

func New(log *zap.Logger, tableConfigRepo ports.TableConfigPersistencePort, token token.Token) *Server {
	server := &Server{Logger: log, Token: token}

	server.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	server.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	kubernetesHandler := NewKubernetesHandler(server)

	server.managmentApp.Get("/healthz/liveness", kubernetesHandler.Liveness)
	server.managmentApp.Get("/healthz/readiness", kubernetesHandler.Readiness)

	tableConfigService := domainService.NewTableConfigService(tableConfigRepo, log, token)
	tableConfigHandler := NewTableConfigHandler(server, appService.NewTableConfigApplicationService(tableConfigService, log))

	eventorV1 := server.clientApp.Group("/api/eventor/v1", tableConfigHandler.fetchUserDataMiddleware)

	eventorV1.Post("/table_config", tableConfigHandler.UpdateTableConfig)
	eventorV1.Get("/table", tableConfigHandler.SeeTable)

	// 404 Handler
	eventorV1.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})

	return server
}

func (handler *Server) Serve() {
	go transaction.RunBitcoinEvent()
	go transaction.RunEthereumEvent()
	go transaction.RunDogecoinEvent()

	go func() {
		err := handler.managmentApp.Listen(":8081")
		handler.Logger.Fatal("error resolving managment server", zap.Error(err))
	}()

	go func() {
		err := handler.clientApp.Listen(":8080")
		handler.Logger.Fatal("error resolving client server", zap.Error(err))
	}()
}
