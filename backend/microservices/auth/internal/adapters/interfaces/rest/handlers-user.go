package rest

import (
	"net/http"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	server         *Server
	userAppService *services.UserApplicationService
}

func NewUserHandler(server *Server, userAppService *services.UserApplicationService) *UserHandler {
	return &UserHandler{server: server, userAppService: userAppService}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {

	var request dto.RegisterRequest
	userAgent := string(c.Request().Header.Peek("User-Agent"))
	if userAgent == "" {
		h.server.Logger.Error("Missing user agent header")
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "header User-Agent", Message: "no user agent header, please provide it"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "body", Message: "Error parsing request body"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	authTokens, err := h.userAppService.Register(&request, userAgent)
	if err != nil {
		if err.Message == "Validation failed" {
			response := dto.ErrorResponse{Errors: err.Details.([]dto.ErrorContent), NeedLogin: false}
			return c.Status(err.StatusCode).JSON(response)
		}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    authTokens.RefreshToken,
		Expires:  time.Now().Add(h.server.Token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := dto.RegisterResponse{
		AccessToken: authTokens.AccessToken,
	}
	return c.Status(http.StatusCreated).JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {

	var request dto.LoginRequest
	userAgent := string(c.Request().Header.Peek("User-Agent"))
	if userAgent == "" {
		h.server.Logger.Error("Missing user agent header")
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "header User-Agent", Message: "no user agent header, please provide it"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "body", Message: "Error parsing request body"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	authTokens, err := h.userAppService.Login(&request, userAgent)
	if err != nil {
		if err.Message == "Validation failed" {
			response := dto.ErrorResponse{Errors: err.Details.([]dto.ErrorContent), NeedLogin: false}
			return c.Status(err.StatusCode).JSON(response)
		}
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    authTokens.RefreshToken,
		Expires:  time.Now().Add(h.server.Token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := dto.LoginResponse{
		AccessToken: authTokens.AccessToken,
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	refreshToken, _ := c.Locals("user-refresh_token").(string)

	err := h.userAppService.Logout(refreshToken)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}
