package rest

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/application/services"
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ArticleHandler struct {
	server            *Server
	articleAppService *services.ArticleApplicationService
}

func NewArticleHandler(server *Server, articleAppService *services.ArticleApplicationService) *ArticleHandler {
	return &ArticleHandler{server: server, articleAppService: articleAppService}
}

func (h *ArticleHandler) NewArticle(c *fiber.Ctx) error {

	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	username, ok := c.Locals("user-username").(string)
	if !ok || username == "" {
		h.server.Logger.Error("Invalid user-username local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	var request dto.NewArticleRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	articleId, err := h.articleAppService.NewArticle(&request, entity.UUID(userId), username)
	if err != nil {
		if err.Message == "Validation failed" {
			response := dto.ErrorResponse{Errors: err.Details.([]dto.ErrorContent), NeedRefresh: false}
			return c.Status(err.StatusCode).JSON(response)
		}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.NewArticleResponse{
		ArticleId: articleId,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (h *ArticleHandler) UpdateArticle(c *fiber.Ctx) error {

	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	var request dto.UpdateArticleRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.articleAppService.UpdateArticle(entity.UUID(userId))
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}

func (h *ArticleHandler) DeleteArticle(c *fiber.Ctx) error {

	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	var request dto.DeleteArticleRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.articleAppService.DeleteArticle(&request)
	if err != nil {
		if err.Message == "Validation failed" {
			response := dto.ErrorResponse{Errors: err.Details.([]dto.ErrorContent), NeedRefresh: false}
			return c.Status(err.StatusCode).JSON(response)
		}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}
