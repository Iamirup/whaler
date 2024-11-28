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
		// response := map[string]string{"error": "no user agent header, please provide it"}
		response := dto.ErrorResponse{Error: "no user agent header, please provide it", NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		// response := map[string]string{"error": "Error parsing request body"}
		response := dto.ErrorResponse{Error: "Error parsing request body", NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err, authTokens := h.userAppService.Register(&request, userAgent)
	if err != nil {
		// response := map[string]string{"error": err.Message}
		response := dto.ErrorResponse{Error: err.Message, NeedLogin: false}
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
		// response := map[string]string{"error": "no user agent header, please provide it"}
		response := dto.ErrorResponse{Error: "no user agent header, please provide it", NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Error(err))
		// response := map[string]string{"error": "Error parsing request body"}
		response := dto.ErrorResponse{Error: "Error parsing request body", NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err, authTokens := h.userAppService.Login(&request, userAgent)
	if err != nil {
		// response := map[string]string{"error": err.Message}
		response := dto.ErrorResponse{Error: err.Message, NeedLogin: false}
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
		// response := map[string]string{"error": err.Message}
		response := dto.ErrorResponse{Error: err.Message, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}
