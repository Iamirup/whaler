package rest

import (
	"github.com/gofiber/fiber/v2"
)

type RefreshTokenHandler struct {
	refreshTokenAppService *services.RefreshTokenApplicationService
}

func (handler *Server) refresh(c *fiber.Ctx) error {

}

func (handler *Server) revoke(c *fiber.Ctx) error {

}
