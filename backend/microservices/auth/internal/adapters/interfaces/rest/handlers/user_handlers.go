package handlers

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

func (h *UserHandler) register(c *fiber.Ctx) error {
	// request := struct{ Username, Password string }{}
	var request dto.RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		h.server.logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := map[string]string{"error": "Error parsing request body"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	/////////////////////////////////////////
	err, authTokens := h.userAppService.Register(c, &request)
	if err != nil {
		response := map[string]string{"error": err.Message}
		return c.Status(err.StatusCode).JSON(response)
	}
	////////////////////////////////////////

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    authTokens.refreshToken,
		Expires:  time.Now().Add(handler.token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := map[string]string{"access_token": authTokens.accessToken}
	return c.Status(http.StatusCreated).JSON(response)
}

func (h *UserHandler) login(c *fiber.Ctx) error {
	// request := struct{ Username, Password string }{}
	var request dto.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		handler.logger.Error("Error parsing request body", zap.Error(err))
		response := map[string]string{"error": "Error parsing request body"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	///////////////////
	err, authTokens := h.userAppService.Login(c, &request)
	if err != nil {
		response := map[string]string{"error": err.Message}
		return c.Status(err.StatusCode).JSON(response)
	}
	///////////////////

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    authTokens.refreshToken,
		Expires:  time.Now().Add(handler.token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := map[string]string{"access_token": authTokens.accessToken}
	return c.Status(http.StatusOK).JSON(response)
}
