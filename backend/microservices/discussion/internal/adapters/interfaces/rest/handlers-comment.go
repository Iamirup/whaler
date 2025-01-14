package rest

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/application/services"
	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/domain/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CommentHandler struct {
	server            *Server
	commentAppService *services.CommentApplicationService
}

func NewCommentHandler(server *Server, commentAppService *services.CommentApplicationService) *CommentHandler {
	return &CommentHandler{server: server, commentAppService: commentAppService}
}

func (h *CommentHandler) NewComment(c *fiber.Ctx) error {

	username, ok := c.Locals("user-username").(string)
	if !ok || username == "" {
		h.server.Logger.Error("Invalid user-username local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	var request dto.NewCommentRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedRefresh: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	commentId, err := h.commentAppService.NewComment(&request, username)
	if err != nil {
		if err.Message == "Validation failed" {
			response := dto.ErrorResponse{Errors: err.Details.([]dto.ErrorContent), NeedRefresh: false}
			return c.Status(err.StatusCode).JSON(response)
		}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.NewCommentResponse{
		CommentId: commentId,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (h *CommentHandler) GetComments(c *fiber.Ctx) error {

	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	cursor := c.Query("cursor")
	limit := c.QueryInt("limit")

	comments, newCursor, err := h.commentAppService.GetComments(entity.UUID(userId), cursor, limit)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedRefresh: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.GetCommentsResponse{
		Comments:  comments,
		NewCursor: newCursor,
	}

	return c.Status(http.StatusOK).JSON(response)
}
