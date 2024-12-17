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
	server            *restServer
	articleAppService *services.ArticleApplicationService
}

func NewArticleHandler(server *restServer, articleAppService *services.ArticleApplicationService) *ArticleHandler {
	return &ArticleHandler{server: server, articleAppService: articleAppService}
}

func (h *ArticleHandler) GetAnArticle(c *fiber.Ctx) error {

	articleId := entity.UUID((c.Params("articleId")))

	article, err := h.articleAppService.GetAnArticle(articleId)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.GetAnArticleResponse{
		Article: *article,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *ArticleHandler) GetAllArticles(c *fiber.Ctx) error {

	cursor := c.Query("cursor")
	limit := c.QueryInt("limit")

	articles, newCursor, err := h.articleAppService.GetAllArticles(cursor, limit)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.GetArticlesResponse{
		Articles:  articles,
		NewCursor: newCursor,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *ArticleHandler) GetMyArticles(c *fiber.Ctx) error {

	authorId, ok := c.Locals("user-id").(string)
	if !ok || authorId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	cursor := c.Query("cursor")
	limit := c.QueryInt("limit")

	articles, newCursor, err := h.articleAppService.GetMyArticles(cursor, limit, entity.UUID(authorId))
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.GetArticlesResponse{
		Articles:  articles,
		NewCursor: newCursor,
	}

	return c.Status(http.StatusOK).JSON(response)
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

	err := h.articleAppService.UpdateArticle(&request, entity.UUID(userId))
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

	err := h.articleAppService.DeleteArticle(&request, entity.UUID(userId))
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

func (h *ArticleHandler) LikeArticle(c *fiber.Ctx) error {

	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	var request dto.LikeArticleRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.articleAppService.LikeArticle(&request, entity.UUID(userId))
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

func (h *ArticleHandler) GetTopAuthors(c *fiber.Ctx) error {

	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	topAuthors, err := h.articleAppService.GetTopAuthors()
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.GetTopAuthorsResponse{
		TopAuthors: topAuthors,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *ArticleHandler) GetPopularArticles(c *fiber.Ctx) error {

	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	articles, err := h.articleAppService.GetPopularArticles()
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.GetPopularArticlesResponse{
		Articles: articles,
	}

	return c.Status(http.StatusOK).JSON(response)
}
