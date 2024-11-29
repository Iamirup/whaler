package rest

import (
	"net/http"
	"strings"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *RefreshTokenHandler) fetchUserDataMiddleware(c *fiber.Ctx) error {
	headerBytes := c.Request().Header.Peek("Authorization")
	header := strings.TrimPrefix(string(headerBytes), "Bearer ")

	if len(header) == 0 {
		h.server.Logger.Error("Missing authorization header")
		// response := map[string]string{"error": "please provide your authentication information"}
		// response := dto.ErrorResponse{Error: "please provide your authentication information", NeedLogin: false}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "please provide your authentication information"}}, NeedLogin: false}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	refreshToken := c.Cookies("refresh_token")
	accessTokenPayload, err := h.server.Token.ExtractTokenData(header)
	if refreshToken == "" {
		h.server.Logger.Error("Missing refresh token")
		// response := map[string]string{"error": "no refresh token header, abnormal activity was detected. please login again"}
		// response := dto.ErrorResponse{Error: "no refresh token header, abnormal activity was detected. please login again", NeedLogin: true}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "no refresh token header, abnormal activity was detected. please login again"}}, NeedLogin: true}
		if err := h.refreshTokenAppService.RevokeAllRefreshTokensById(accessTokenPayload.Id); err != nil {
			h.server.Logger.Error("something went wrong")
			// response := map[string]string{"error": "Something went wrong! please try again later"}
			// response := dto.ErrorResponse{Error: "Something went wrong! please try again later", NeedLogin: false}
			response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Something went wrong! please try again later"}}, NeedLogin: false}
			return c.Status(http.StatusInternalServerError).JSON(response)
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	if err != nil {
		if err.Error() == "error token has expired" {
			err = h.server.Token.ValidateRefreshToken(refreshToken)
			if err != nil {
				h.server.Logger.Error("Invalid refresh token", zap.Error(err))
				// response := map[string]string{"error": "invalid refresh token, abnormal activity was detected. please login again"}
				// response := dto.ErrorResponse{Error: "invalid refresh token, abnormal activity was detected. please login again", NeedLogin: true}
				response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "invalid refresh token, abnormal activity was detected. please login again"}}, NeedLogin: true}
				if err := h.refreshTokenAppService.RevokeAllRefreshTokensById(accessTokenPayload.Id); err != nil {
					h.server.Logger.Error("something went wrong")
					// response := map[string]string{"error": "Something went wrong! please try again later"}
					// response := dto.ErrorResponse{Error: "Something went wrong! please try again later", NeedLogin: false}
					response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Something went wrong! please try again later"}}, NeedLogin: false}
					return c.Status(http.StatusInternalServerError).JSON(response)
				}
				return c.Status(http.StatusUnauthorized).JSON(response)
			}

			if err := h.refreshTokenAppService.CheckRefreshTokenInDBById(accessTokenPayload.Id, refreshToken); err != nil {
				// response := map[string]string{"error": err.Message}
				// response := dto.ErrorResponse{Error: err.Message, NeedLogin: false}
				response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
				return c.Status(err.StatusCode).JSON(response)
			}

			isAdmin, err := h.refreshTokenAppService.CheckIfIsAdmin(accessTokenPayload.Id)
			if err != nil {
				h.server.Logger.Error("something went wrong")
				// response := map[string]string{"error": "Something went wrong! please try again later"}
				// response := dto.ErrorResponse{Error: "Something went wrong! please try again later", NeedLogin: false}
				response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Something went wrong! please try again later"}}, NeedLogin: false}
				return c.Status(http.StatusInternalServerError).JSON(response)
			}

			newAccessToken, errs := h.server.Token.CreateTokenString(accessTokenPayload.Id, accessTokenPayload.Username, isAdmin)
			if errs != nil {
				h.server.Logger.Error("Failed to create new access token", zap.Error(errs))
				// response := map[string]string{"error": "failed to create new access token, abnormal activity was detected. please login again"}
				// response := dto.ErrorResponse{Error: "failed to create new access token, abnormal activity was detected. please login again", NeedLogin: true}
				response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "failed to create new access token, abnormal activity was detected. please login again"}}, NeedLogin: true}
				if err := h.refreshTokenAppService.RevokeAllRefreshTokensById(accessTokenPayload.Id); err != nil {
					h.server.Logger.Error("something went wrong")
					// response := map[string]string{"error": "Something went wrong! please try again later"}
					// response := dto.ErrorResponse{Error: "Something went wrong! please try again later", NeedLogin: false}
					response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Something went wrong! please try again later"}}, NeedLogin: false}
					return c.Status(http.StatusInternalServerError).JSON(response)
				}
				return c.Status(http.StatusInternalServerError).JSON(response)
			}

			c.Set("Authorization", "Bearer "+newAccessToken)

		} else {
			h.server.Logger.Error("Something is wrong with access token")
			// response := map[string]string{"error": "invalid token header, abnormal activity was detected. please login again"}
			// response := dto.ErrorResponse{Error: "invalid token header, abnormal activity was detected. please login again", NeedLogin: true}
			response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "invalid token header, abnormal activity was detected. please login again"}}, NeedLogin: true}
			if err := h.refreshTokenAppService.RevokeAllRefreshTokensById(accessTokenPayload.Id); err != nil {
				h.server.Logger.Error("something went wrong")
				// response := map[string]string{"error": "Something went wrong! please try again later"}
				// response := dto.ErrorResponse{Error: "Something went wrong! please try again later", NeedLogin: false}
				response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Something went wrong! please try again later"}}, NeedLogin: false}
				return c.Status(http.StatusInternalServerError).JSON(response)
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
	}

	c.Locals("user-id", accessTokenPayload.Id)
	c.Locals("user-username", accessTokenPayload.Username)
	c.Locals("user-refresh_token", refreshToken)
	return c.Next()
}
