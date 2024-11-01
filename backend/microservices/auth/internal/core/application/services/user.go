package services

import (
	"context"
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"

	api "github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
)

// VoucherApplicationService provides application logic for vouchers.
type UserApplicationService struct {
	domainService ports.UserServicePort
}

// NewVoucherApplicationService creates a new instance of VoucherApplicationService.
func NewUserApplicationService(domainService ports.UserServicePort) *UserApplicationService {
	return &UserApplicationService{
		domainService: domainService,
	}
}

// RedeemVoucher handles the redemption process of a voucher and interacts with the domain services.
func (s *UserApplicationService) Register(ctx context.Context, request *api.RegisterRequest) (*serr.ServiceError, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return &serr.ServiceError{Message: "no valid request", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	err, authTokens := s.domainService.Register(ctx, request)
	if err != nil {
		return err, entity.AuthTokens{}
	}

	return nil, authTokens
}

// CreateVoucher create a new voucher
func (s *UserApplicationService) Login(ctx context.Context, request *api.LoginRequest) (*serr.ServiceError, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return &serr.ServiceError{Message: "no valid request", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}
	return s.domainService.Login(ctx, request.Code, request.UsageLimit)
}
