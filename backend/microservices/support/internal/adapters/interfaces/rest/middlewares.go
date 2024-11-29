package rest

import (
	"net/http"
	"strings"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/adapters/interfaces/rest/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TicketHandler) fetchUserDataMiddleware(c *fiber.Ctx) error {
	headerBytes := c.Request().Header.Peek("Authorization")
	header := strings.TrimPrefix(string(headerBytes), "Bearer ")

	if len(header) == 0 {
		h.server.Logger.Error("Missing authorization header")
		response := map[string]string{"error": "please provide your authentication information"}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	accessTokenPayload, err := h.server.Token.ExtractTokenData(header)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "need refresh"}}, NeedRefresh: true}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	c.Locals("user-id", accessTokenPayload.Id)
	c.Locals("user-username", accessTokenPayload.Username)
	c.Locals("user-is_admin", accessTokenPayload.IsAdmin)
	return c.Next()
}
