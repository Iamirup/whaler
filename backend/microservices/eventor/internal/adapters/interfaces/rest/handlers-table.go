package rest

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type TableConfigHandler struct {
	server                *Server
	tableConfigAppService *services.TableConfigApplicationService
}

func NewTableConfigHandler(server *Server, tableConfigAppService *services.TableConfigApplicationService) *TableConfigHandler {
	return &TableConfigHandler{server: server, tableConfigAppService: tableConfigAppService}
}

func (h *TableConfigHandler) UpdateTableConfig(c *fiber.Ctx) error {

	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	var request dto.UpdateTableConfigRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// err := h.tableConfigAppService.UpdateTableConfig(&request)
	// if err != nil {
	// 	if err.Message == "Validation failed" {
	// 		response := dto.ErrorResponse{Errors: err.Details.([]dto.ErrorContent), NeedRefresh: false}
	// 		return c.Status(err.StatusCode).JSON(response)
	// 	}
	// 	response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
	// 	return c.Status(err.StatusCode).JSON(response)
	// }

	return c.SendStatus(http.StatusCreated)
}

func (h *TableConfigHandler) SeeTable(c *fiber.Ctx) error {

	cursor := c.Query("cursor")
	limit := c.QueryInt("limit")

	cryptocurrency := c.Query("cryptocurrency")
	minAmount := c.QueryFloat("minAmount")
	age := c.QueryInt("age")

	table, newCursor, err := h.tableConfigAppService.SeeTable(cryptocurrency, minAmount, age, cursor, limit)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.SeeTableResponse{
		Table:     table,
		NewCursor: newCursor,
	}

	return c.Status(http.StatusOK).JSON(response)
}
