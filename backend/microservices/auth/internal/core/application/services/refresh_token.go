package services

import (
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
)

type RefreshTokenApplicationService struct {
	domainService ports.RefreshTokenServicePort
}

func NewRefreshTokenApplicationService(domainService ports.RefreshTokenServicePort) *RefreshTokenApplicationService {
	return &RefreshTokenApplicationService{
		domainService: domainService,
	}
}

func (s *RefreshTokenApplicationService) GetAndCheckRefreshTokenById(userId, refreshToken string) *serr.ServiceError {
	return s.domainService.GetAndCheckRefreshTokenById(userId, refreshToken)
}
