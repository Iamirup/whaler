package services

import (
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
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

func (s *RefreshTokenApplicationService) GetRefreshTokenById(userId string) (*entity.RefreshToken, *serr.ServiceError) {
	return s.domainService.GetRefreshTokenById(userId)
}
