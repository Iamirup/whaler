package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/application/ports"
	appService "github.com/Iamirup/whaler/backend/microservices/support/internal/core/application/services"
	domainService "github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/services"
	"github.com/Iamirup/whaler/backend/microservices/support/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
	Token  token.Token

	managmentApp *fiber.App
	clientApp    *fiber.App
}

func New(log *zap.Logger, ticketRepo ports.TicketPersistencePort, token token.Token) *Server {
	server := &Server{Logger: log, Token: token}

	server.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	server.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	kubernetesHandler := NewKubernetesHandler(server)

	server.managmentApp.Get("/healthz/liveness", kubernetesHandler.Liveness)
	server.managmentApp.Get("/healthz/readiness", kubernetesHandler.Readiness)

	ticketService := domainService.NewTicketService(ticketRepo, log, token)
	ticketHandler := NewTicketHandler(server, appService.NewTicketApplicationService(ticketService, log))

	supportV1 := server.clientApp.Group("/api/support/v1", ticketHandler.fetchUserDataMiddleware)

	supportV1.Post("/ticket/new", ticketHandler.NewTicket)       // for users
	supportV1.Get("/tickets/me", ticketHandler.MyTickets)        // for users
	supportV1.Post("/ticket/reply", ticketHandler.ReplyToTicket) // for admin
	supportV1.Get("/tickets/all", ticketHandler.AllTicket)       // for admin

	// 404 Handler
	supportV1.Use(func(c *fiber.Ctx) error {
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
