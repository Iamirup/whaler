package handlers

import (
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/services"
	"github.com/gofiber/fiber/v2"
)

type RefreshTokenHandler struct {
	server                 *rest.Server
	refreshTokenAppService *services.RefreshTokenApplicationService
}

func (h *RefreshTokenHandler) refresh(c *fiber.Ctx) error {

}

func (h *RefreshTokenHandler) revoke(c *fiber.Ctx) error {

}
