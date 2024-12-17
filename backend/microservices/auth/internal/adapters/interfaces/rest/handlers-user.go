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
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "body", Message: "Error parsing request body"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	authTokens, err := h.userAppService.Register(&request)
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
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    authTokens.AccessToken,
		Expires:  time.Now().Add(h.server.Token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.SendStatus(http.StatusCreated)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {

	var request dto.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "body", Message: "Error parsing request body"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	possibleRefreshToken := c.Cookies("refresh_token")

	authTokens, err := h.userAppService.Login(&request, possibleRefreshToken)
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
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    authTokens.AccessToken,
		Expires:  time.Now().Add(h.server.Token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.SendStatus(http.StatusOK)
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

func (h *UserHandler) IsAdmin(c *fiber.Ctx) error {
	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	return c.SendStatus(http.StatusOK)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	var request dto.DeleteUserRequest
	if err := c.BodyParser(&request); err != nil {
		h.server.Logger.Error("Error parsing request body", zap.Any("request", request), zap.Error(err))
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: "Error parsing request body"}}, NeedLogin: false}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	err := h.userAppService.DeleteUser(&request)
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	return c.SendStatus(http.StatusOK)
}

func (h *UserHandler) GetOnlineUsers(c *fiber.Ctx) error {
	isAdmin, ok := c.Locals("user-is_admin").(bool)
	if !ok {
		h.server.Logger.Error("Invalid user-is_admin local")
		return c.SendStatus(http.StatusInternalServerError)
	} else if !isAdmin {
		h.server.Logger.Error("Forbidden access")
		return c.SendStatus(http.StatusForbidden)
	}

	onlineUsers, err := h.userAppService.GetOnlineUsers()
	if err != nil {
		response := dto.ErrorResponse{Errors: []dto.ErrorContent{{Field: "_", Message: err.Message}}, NeedLogin: false}
		return c.Status(err.StatusCode).JSON(response)
	}

	response := dto.GetOnlineUsersResponse{
		OnlineUsers: onlineUsers,
	}

	return c.Status(http.StatusOK).JSON(response)
}
