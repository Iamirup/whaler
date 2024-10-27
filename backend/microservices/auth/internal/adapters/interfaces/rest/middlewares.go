package rest

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (handler *Server) fetchUserId(c *fiber.Ctx) error {
	headerBytes := c.Request().Header.Peek("Authorization")
	header := strings.TrimPrefix(string(headerBytes), "Bearer ")

	if len(header) == 0 {
		handler.logger.Error("Missing authorization header")
		response := map[string]string{"error": "please provide your authentication information"}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	var id string
	err := handler.token.ExtractTokenData(header, &id)
	if err != nil {
		// Attempt to use refresh token if access token is invalid or expired
		refreshToken := c.Cookies("refresh_token")
		if refreshToken == "" {
			handler.logger.Error("Missing refresh token")
			response := map[string]string{"error": "invalid token header, please login again"}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		if err.Error() == "error token has expired" {
			// from db
			err = handler.token.ValidateRefreshToken(refreshToken)
			if err != nil {
				handler.logger.Error("Invalid refresh token", zap.Error(err))
				response := map[string]string{"error": "invalid refresh token, please login again"}
				return c.Status(http.StatusUnauthorized).JSON(response)
			}

			DBrefreshToken, err := handler.repository.GetRefreshTokenById(id)
			if err != nil || DBrefreshToken == nil {
				handler.logger.Error("Error invalid refresh token returned", zap.Error(err))
				response := map[string]string{"error": "invalid refresh token, please login again"}
				return c.Status(http.StatusInternalServerError).JSON(response)
			}

			if refreshToken != DBrefreshToken.Token {
				handler.logger.Error("Invalid refresh token", zap.Error(err))
				response := map[string]string{"error": "invalid refresh token, please login again"}
				return c.Status(http.StatusUnauthorized).JSON(response)
			}

			// Generate new access token
			newAccessToken, err := handler.token.CreateTokenString(id)
			if err != nil {
				handler.logger.Error("Failed to create new access token", zap.Error(err))
				response := map[string]string{"error": "failed to create new access token, please login again"}
				return c.Status(http.StatusInternalServerError).JSON(response)
			}

			// Set new access token in response header
			c.Set("Authorization", "Bearer "+newAccessToken)

		} else {
			handler.logger.Error("Something is wrong with access token")
			response := map[string]string{"error": "invalid token header, please login again"}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
	}

	c.Locals("user-id", id)
	return c.Next()
}
