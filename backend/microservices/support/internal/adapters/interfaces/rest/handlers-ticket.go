package rest

import (
	"net/http"
	"strconv"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/application/services"
	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type TicketHandler struct {
	server           *Server
	ticketAppService *services.TicketApplicationService
}

func NewTicketHandler(server *Server, ticketAppService *services.TicketApplicationService) *TicketHandler {
	return &TicketHandler{server: server, ticketAppService: ticketAppService}
}

func (h *TicketHandler) NewTicket(c *fiber.Ctx) error {

	userId, ok := c.Locals("user-id").(entity.UUID)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	username, ok := c.Locals("user-username").(string)
	if !ok || username == "" {
		h.server.Logger.Error("Invalid user-username local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	var request dto.NewTicketRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		// response := map[string]string{"error": "Error parsing request body"}
		response := dto.ErrorResponse{Error: "Error parsing request body", NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	ticketId, err := h.ticketAppService.NewTicket(&request, userId, username)
	if err != nil {
		// response := map[string]string{"error": err.Message}
		response := dto.ErrorResponse{Error: err.Message, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.NewTicketResponse{
		TicketId: ticketId,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (h *TicketHandler) MyTickets(c *fiber.Ctx) error {

	userId, ok := c.Locals("user-id").(entity.UUID)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	cursor := c.Query("cursor")
	limit, _ := strconv.Atoi(c.Query("limit"))

	tickets, newCursor, err := h.ticketAppService.MyTickets(userId, cursor, limit)
	if err != nil {
		// response := map[string]string{"error": err.Message}
		response := dto.ErrorResponse{Error: err.Message, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.MyTicketsResponse{
		Tickets:   tickets,
		NewCursor: newCursor,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *TicketHandler) ReplyToTicket(c *fiber.Ctx) error {

	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if isAdmin == false {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	var request dto.ReplyToTicketRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		// response := map[string]string{"error": "Error parsing request body"}
		response := dto.ErrorResponse{Error: "Error parsing request body", NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.ticketAppService.ReplyToTicket(&request)
	if err != nil {
		// response := map[string]string{"error": err.Message}
		response := dto.ErrorResponse{Error: err.Message, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusCreated)
}

func (h *TicketHandler) AllTicket(c *fiber.Ctx) error {

	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if isAdmin == false {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	cursor := c.Query("cursor")
	limit, _ := strconv.Atoi(c.Query("limit"))

	tickets, newCursor, err := h.ticketAppService.AllTicket(cursor, limit)
	if err != nil {
		// response := map[string]string{"error": err.Message}
		response := dto.ErrorResponse{Error: err.Message, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.AllTicketResponse{
		Tickets:   tickets,
		NewCursor: newCursor,
	}

	return c.Status(http.StatusOK).JSON(response)
}
