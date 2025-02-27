package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	appService "github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/services"
	domainService "github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/services"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
	Token  token.Token

	managmentApp *fiber.App
	clientApp    *fiber.App
}

func New(log *zap.Logger, userRepo ports.UserPersistencePort, adminRepo ports.AdminPersistencePort, refreshTokenRepo ports.RefreshTokenPersistencePort, token token.Token) *Server {
	server := &Server{Logger: log, Token: token}

	server.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	server.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	kubernetesHandler := NewKubernetesHandler(server)

	server.managmentApp.Get("/healthz/liveness", kubernetesHandler.Liveness)
	server.managmentApp.Get("/healthz/readiness", kubernetesHandler.Readiness)

	authV1 := server.clientApp.Group("/api/auth/v1")

	userService := domainService.NewUserService(userRepo, refreshTokenRepo, log, token)
	userHandler := NewUserHandler(server, appService.NewUserApplicationService(userService, log))

	adminService := domainService.NewAdminService(adminRepo, log, token)
	adminHandler := NewAdminHandler(server, appService.NewAdminApplicationService(adminService, log))

	refreshTokenService := domainService.NewRefreshTokenService(refreshTokenRepo, log, token)
	refreshTokenHandler := NewRefreshTokenHandler(server, appService.NewRefreshTokenApplicationService(refreshTokenService))

	authV1.Post("/register", userHandler.Register)
	authV1.Post("/login", userHandler.Login)
	authV1.Post("/logout", refreshTokenHandler.fetchUserDataMiddleware, userHandler.Logout)
	authV1.Get("/refresh", refreshTokenHandler.fetchUserDataMiddleware, refreshTokenHandler.Refresh)
	authV1.Get("/is_admin", refreshTokenHandler.fetchUserDataMiddleware, userHandler.IsAdmin)
	authV1.Delete("/user", refreshTokenHandler.fetchUserDataMiddleware, userHandler.DeleteUser)     // for admin
	authV1.Post("/admin", refreshTokenHandler.fetchUserDataMiddleware, adminHandler.AddAdmin)       // for admin
	authV1.Delete("/admin", refreshTokenHandler.fetchUserDataMiddleware, adminHandler.DeleteAdmin)  // for admin
	authV1.Get("/onlines", refreshTokenHandler.fetchUserDataMiddleware, userHandler.GetOnlineUsers) // for admin

	// 404 Handler
	authV1.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})

	return server
}

func (handler *Server) Serve() {
	go func() {
		err := handler.managmentApp.Listen(":8081")
		handler.Logger.Fatal("error resolving managment server", zap.Error(err))
	}()

	go func() {
		err := handler.clientApp.Listen(":8080")
		handler.Logger.Fatal("error resolving client server", zap.Error(err))
	}()
}
