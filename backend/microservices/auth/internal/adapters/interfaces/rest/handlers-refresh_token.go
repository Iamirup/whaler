package rest

import (
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
	return nil
}

func (h *RefreshTokenHandler) Revoke(c *fiber.Ctx) error {
	return nil
}
