package services

import (
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RefreshTokenService struct {
	refreshTokenPersistencePort ports.RefreshTokenPersistencePort
	logger                      *zap.Logger
	token                       token.Token
}

func NewRefreshTokenService(
	refreshTokenPersistencePort ports.RefreshTokenPersistencePort,
	logger *zap.Logger, token token.Token) ports.RefreshTokenServicePort {

	return &RefreshTokenService{
		refreshTokenPersistencePort: refreshTokenPersistencePort,
		logger:                      logger,
		token:                       token,
	}
}

func (s *RefreshTokenService) Refresh(ctx *fiber.Ctx, something string) error {
	return nil
}

func (s *RefreshTokenService) Revoke(ctx *fiber.Ctx, something string) error {
	return nil
}
