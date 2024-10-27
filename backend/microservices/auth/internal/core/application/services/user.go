package services

import (
	"context"
	"voucher/internal/core/application/ports"
	api "voucher/internal/interfaces/api/dto"
	"voucher/pkg/serr"
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
func (s *UserApplicationService) Register(ctx context.Context, request *api.RegisterRequest) (error, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return err
	}

	err, authTokens := s.domainService.Register(ctx, request)
	if err != nil {
		return err
	}

	return nil, authTokens
}

// CreateVoucher create a new voucher
func (s *UserApplicationService) Login(ctx context.Context, request *api.CreateVoucherRequest) (error, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return serr.ValidationErr("VoucherApplicationService.CreateVoucher",
			err.Error(), serr.ErrInvalidInput)
	}
	return s.domainService.Login(ctx, request.Code, request.UsageLimit)
}
