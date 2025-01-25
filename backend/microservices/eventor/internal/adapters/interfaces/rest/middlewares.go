package rest

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/adapters/interfaces/rest/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TableConfigHandler) fetchUserDataMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		h.server.Logger.Error("Missing access token")
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "please provide your authentication information"}}, NeedRefresh: true}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	accessTokenPayload, err := h.server.Token.ExtractTokenData(accessToken)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "need refresh"}}, NeedRefresh: true}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	c.Locals("user-is_admin", accessTokenPayload.IsAdmin)
	return c.Next()
}
