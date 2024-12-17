package rest

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AdminHandler struct {
	server          *Server
	adminAppService *services.AdminApplicationService
}

func NewAdminHandler(server *Server, adminAppService *services.AdminApplicationService) *AdminHandler {
	return &AdminHandler{server: server, adminAppService: adminAppService}
}

func (h *AdminHandler) AddAdmin(c *fiber.Ctx) error {
	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	var request dto.AddAdminRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.adminAppService.AddAdmin(&request)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}

func (h *AdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	var request dto.DeleteAdminRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.adminAppService.DeleteAdmin(&request)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}
