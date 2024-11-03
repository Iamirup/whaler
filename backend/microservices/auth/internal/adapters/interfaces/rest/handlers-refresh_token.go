package rest

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/services"

	"github.com/gofiber/fiber/v2"
)

type RefreshTokenHandler struct {
	server                 *Server
	refreshTokenAppService *services.RefreshTokenApplicationService
}

func NewRefreshTokenHandler(server *Server, refreshTokenAppService *services.RefreshTokenApplicationService) *RefreshTokenHandler {
	return &RefreshTokenHandler{server: server, refreshTokenAppService: refreshTokenAppService}
}

func (h *RefreshTokenHandler) Refresh(c *fiber.Ctx) error {
	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		h.server.Logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}
