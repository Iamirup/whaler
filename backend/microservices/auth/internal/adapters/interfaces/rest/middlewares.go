package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *UserHandler) fetchUserRefreshTokenMiddleware(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")

	fmt.Println(refreshToken)

	if refreshToken == "" {
		h.server.Logger.Error("Missing refresh token")
		response := map[string]string{"error": "no refresh token header, please login again"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.server.Token.ValidateRefreshToken(refreshToken)
	if err != nil {
		h.server.Logger.Error("Invalid refresh token", zap.Error(err))
		response := map[string]string{"error": "invalid refresh token, please login again"}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	fmt.Println(refreshToken)

	c.Locals("user-refresh_token", refreshToken)
	return c.Next()
}

func (h *RefreshTokenHandler) fetchUserIdMiddleware(c *fiber.Ctx) error {
	headerBytes := c.Request().Header.Peek("Authorization")
	header := strings.TrimPrefix(string(headerBytes), "Bearer ")

	if len(header) == 0 {
		h.server.Logger.Error("Missing authorization header")
		response := map[string]string{"error": "please provide your authentication information"}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	accessTokenPayload, err := h.server.Token.ExtractTokenData(header)
	id, username := accessTokenPayload.Id, accessTokenPayload.Username
	if err != nil {
		// Attempt to use refresh token if access token is invalid or expired
		refreshToken := c.Cookies("refresh_token")
		if refreshToken == "" {
			h.server.Logger.Error("Missing refresh token")
			response := map[string]string{"error": "no refresh token header, please login again"}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		if err.Error() == "error token has expired" {
			err = h.server.Token.ValidateRefreshToken(refreshToken)
			if err != nil {
				h.server.Logger.Error("Invalid refresh token", zap.Error(err))
				response := map[string]string{"error": "invalid refresh token, please login again"}
				return c.Status(http.StatusUnauthorized).JSON(response)
			}

			if err := h.refreshTokenAppService.GetAndCheckRefreshTokenById(id, refreshToken); err != nil {
				response := map[string]string{"error": err.Message}
				return c.Status(err.StatusCode).JSON(response)
			}

			// if !entity.CheckTokenHash(refreshToken, DBrefreshTokenEntity.Token) {
			// 	h.server.Logger.Error("Invalid refresh token")
			// 	response := map[string]string{"error": "invalid refresh token, please login again"}
			// 	return c.Status(http.StatusUnauthorized).JSON(response)
			// }

			// Generate new access token
			newAccessToken, errs := h.server.Token.CreateTokenString(id, username)
			if errs != nil {
				h.server.Logger.Error("Failed to create new access token", zap.Error(errs))
				response := map[string]string{"error": "failed to create new access token, please login again"}
				return c.Status(http.StatusInternalServerError).JSON(response)
			}

			// Set new access token in response header
			c.Set("Authorization", "Bearer "+newAccessToken)

		} else {
			h.server.Logger.Error("Something is wrong with access token")
			response := map[string]string{"error": "invalid token header, please login again"}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
	}

	c.Locals("user-id", id)
	c.Locals("user-username", username)
	return c.Next()
}
