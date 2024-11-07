package services

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
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

func (s *RefreshTokenService) GetAndCheckRefreshTokenById(userId, refreshToken string) *serr.ServiceError {
	if err := s.refreshTokenPersistencePort.GetAndCheckRefreshTokenById(userId, refreshToken); err != nil {
		s.logger.Error("Error invalid refresh token returned", zap.Error(err))
		if err := s.refreshTokenPersistencePort.RevokeAllRefreshTokensById(userId); err != nil {
			s.logger.Error("something went wrong")
			return &serr.ServiceError{Message: "Something went wrong! please try again later", StatusCode: http.StatusInternalServerError}
		}
		return &serr.ServiceError{Message: "invalid refresh token, abnormal activity was detected. please login again", StatusCode: http.StatusInternalServerError}
	}

	return nil
}

func (s *RefreshTokenService) RevokeAllRefreshTokensById(userId string) error {
	return s.refreshTokenPersistencePort.RevokeAllRefreshTokensById(userId)
}
