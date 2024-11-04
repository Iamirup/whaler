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

	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := map[string]string{"error": "Error parsing request body"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	/////////////////////////////////////////
	err, authTokens := h.userAppService.Register(&request)
	if err != nil {
		response := map[string]string{"error": err.Message}
		return c.Status(err.StatusCode).JSON(response)
	}
	/////////////////////////////////////////

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    authTokens.RefreshToken,
		Expires:  time.Now().Add(h.server.Token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := map[string]string{"access_token": authTokens.AccessToken}
	return c.Status(http.StatusCreated).JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {

	var request dto.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Error(err))
		response := map[string]string{"error": "Error parsing request body"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	///////////////////
	err, authTokens := h.userAppService.Login(&request)
	if err != nil {
		response := map[string]string{"error": err.Message}
		return c.Status(err.StatusCode).JSON(response)
	}
	///////////////////

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    authTokens.RefreshToken,
		Expires:  time.Now().Add(h.server.Token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := map[string]string{"access_token": authTokens.AccessToken}
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	refreshToken, ok := c.Locals("user-refresh_token").(string)
	if !ok || refreshToken == "" {
		h.server.Logger.Error("Invalid user-refresh_token local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	err := h.userAppService.Logout(refreshToken)
	if err != nil {
		response := map[string]string{"error": err.Message}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}
