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

func (s *RefreshTokenApplicationService) CheckRefreshTokenInDBById(userId string) *serr.ServiceError {
	return s.domainService.CheckRefreshTokenInDBById(userId)
}

func (s *RefreshTokenApplicationService) RevokeAllRefreshTokensById(userId string) error {
	return s.domainService.RevokeAllRefreshTokensById(userId)
}

func (s *RefreshTokenApplicationService) CheckIfIsAdmin(userId string) (bool, error) {
	return s.domainService.CheckIfIsAdmin(userId)
}
