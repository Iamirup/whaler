package services

import (
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"github.com/gofiber/fiber/v2"

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
func (s *UserApplicationService) Register(ctx *fiber.Ctx, request *api.RegisterRequest) (*serr.ServiceError, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return &serr.ServiceError{Message: "no valid request", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	return s.domainService.Register(ctx, request.Username, request.Password)
}

// CreateVoucher create a new voucher
func (s *UserApplicationService) Login(ctx *fiber.Ctx, request *api.LoginRequest) (*serr.ServiceError, entity.AuthTokens) {
	err := request.Validate()
	if err != nil {
		return &serr.ServiceError{Message: "no valid request", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	return s.domainService.Login(ctx, request.Username, request.Password)
}
