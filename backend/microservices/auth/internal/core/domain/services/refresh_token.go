package services

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"
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

func (s *RefreshTokenService) GetRefreshTokenById(userId string) (*entity.RefreshToken, *serr.ServiceError) {
	refreshToken, err := s.refreshTokenPersistencePort.GetRefreshTokenById(userId)
	if err != nil {
		s.logger.Error("Error invalid refresh token returned", zap.Error(err))
		return nil, &serr.ServiceError{Message: "invalid refresh token, please login again", StatusCode: http.StatusInternalServerError}
	}

	return refreshToken, nil
}
