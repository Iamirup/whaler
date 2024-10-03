package http

import (
	"fmt"
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
		response := "please provide your authentication information"
		return c.Status(http.StatusUnauthorized).SendString(response)
	}

	var id uint64
	err := handler.token.ExtractTokenData(header, &id)
	if err != nil {
		// Attempt to use refresh token if access token is invalid or expired
		refreshToken := c.Cookies("refresh_token")
		if refreshToken == "" {
			handler.logger.Error("Missing refresh token")
			response := "invalid token header, please login again"
			return c.Status(http.StatusBadRequest).SendString(response)
		}

		fmt.Println(err.Error())

		if strings.Contains(err.Error(), "token is expired") {
			// from db
			err = handler.token.ValidateRefreshToken(refreshToken)
			if err != nil {
				handler.logger.Error("Invalid refresh token", zap.Error(err))
				response := "invalid refresh token, please login again"
				return c.Status(http.StatusUnauthorized).SendString(response)
			}

			DBrefreshToken, err := handler.repository.GetRefreshTokenById(id)
			if err != nil || DBrefreshToken == nil {
				handler.logger.Error("Error invalid refresh token returned", zap.Error(err))
				response := "invalid refresh token, please login again"
				return c.Status(http.StatusInternalServerError).SendString(response)
			}

			if refreshToken != DBrefreshToken.Token {
				handler.logger.Error("Invalid refresh token", zap.Error(err))
				response := "invalid refresh token, please login again"
				return c.Status(http.StatusUnauthorized).SendString(response)
			}

			// Generate new access token
			newAccessToken, err := handler.token.CreateTokenString(id)
			if err != nil {
				handler.logger.Error("Failed to create new access token", zap.Error(err))
				response := "failed to create new access token, please login again"
				return c.Status(http.StatusInternalServerError).SendString(response)
			}

			// Set new access token in response header
			c.Set("Authorization", "Bearer "+newAccessToken)

		} else {
			handler.logger.Error("Something is wrong with access token")
			response := "invalid token header, please login again"
			return c.Status(http.StatusBadRequest).SendString(response)
		}
	}

	c.Locals("user-id", id)
	return c.Next()
}
