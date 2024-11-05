package services

import (
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
)

type UserApplicationService struct {
	domainService ports.UserServicePort
	logger        *zap.Logger
}

func NewUserApplicationService(domainService ports.UserServicePort, logger *zap.Logger) *UserApplicationService {
	return &UserApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *UserApplicationService) Register(request *api.RegisterRequest, userAgent string) (*serr.ServiceError, entity.AuthTokens) {
	if err := request.Validate(); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return &serr.ServiceError{Message: err.Error(), StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	return s.domainService.Register(request.Email, request.Username, request.Password, userAgent)
}

func (s *UserApplicationService) Login(request *api.LoginRequest, userAgent string) (*serr.ServiceError, entity.AuthTokens) {
	if err := request.Validate(); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return &serr.ServiceError{Message: err.Error(), StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	return s.domainService.Login(request.Email, request.Username, request.Password, userAgent)
}

func (s *UserApplicationService) Logout(refreshToken string) *serr.ServiceError {
	return s.domainService.Logout(refreshToken)
}
