package services

import (
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"

	api "github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
)

type UserApplicationService struct {
	domainService ports.UserServicePort
}

func NewUserApplicationService(domainService ports.UserServicePort) *UserApplicationService {
	return &UserApplicationService{
		domainService: domainService,
	}
}

// RedeemVoucher handles the redemption process of a voucher and interacts with the domain services.
func (s *UserApplicationService) Register(request *api.RegisterRequest) (*serr.ServiceError, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return &serr.ServiceError{Message: "no valid request", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	if request.Password != request.ConfirmPassword {
		return &serr.ServiceError{Message: "password and confirm password are not equal", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	return s.domainService.Register(request.Email, request.Username, request.Password)
}

// CreateVoucher create a new voucher
func (s *UserApplicationService) Login(request *api.LoginRequest) (*serr.ServiceError, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return &serr.ServiceError{Message: "no valid request", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	return s.domainService.Login(request.Username, request.Password)
}

func (s *UserApplicationService) Logout(refreshToken string) *serr.ServiceError {
	return s.domainService.Logout(refreshToken)
}
