package rest

import (
	"net/http"
	"strconv"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/application/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type NewsHandler struct {
	server         *Server
	newsAppService *services.NewsApplicationService
}

func NewNewsHandler(server *Server, newsAppService *services.NewsApplicationService) *NewsHandler {
	return &NewsHandler{server: server, newsAppService: newsAppService}
}

func (h *NewsHandler) AddNews(c *fiber.Ctx) error {

	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	var request dto.AddNewsRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	newsId, err := h.newsAppService.AddNews(&request)
	if err != nil {
		if err.Message == "Validation failed" {
			response := dto.ErrorResponse{Errors: err.Details.([]dto.ErrorContent), NeedRefresh: false}
			return c.Status(err.StatusCode).JSON(response)
		}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.AddNewsResponse{
		NewsId: newsId,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (h *NewsHandler) SeeNews(c *fiber.Ctx) error {

	cursor := c.Query("cursor")
	limit, _ := strconv.Atoi(c.Query("limit"))

	news, newCursor, err := h.newsAppService.SeeNews(cursor, limit)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.SeeNewsResponse{
		News:      news,
		NewCursor: newCursor,
	}

	return c.Status(http.StatusOK).JSON(response)
}
