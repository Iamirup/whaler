package services

import (
	"fmt"
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	"github.com/iancoleman/strcase"
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

func (s *UserApplicationService) Register(request *api.RegisterRequest, userAgent string) (entity.AuthTokens, *serr.ServiceError) {

	if err := request.Validate(); err != nil {
		var validationErrors []api.ErrorContent
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("field '%s' is required", strcase.ToSnake(err.Field()))
			case "email":
				message = fmt.Sprintf("field '%s' must be a valid email address", strcase.ToSnake(err.Field()))
			case "username":
				message = fmt.Sprintf("field '%s' must be a valid username", strcase.ToSnake(err.Field()))
			case "strong_password":
				message = fmt.Sprintf("field '%s' must be a strong password", strcase.ToSnake(err.Field()))
			case "eqfield":
				message = fmt.Sprintf("field '%s' must be equal to '%s'", strcase.ToSnake(err.Field()), strcase.ToSnake(err.Param()))
			default:
				message = fmt.Sprintf("field '%s' failed validation on the '%s' tag", strcase.ToSnake(err.Field()), err.Tag())
			}

			validationErrors = append(validationErrors, api.ErrorContent{
				Field:   strcase.ToSnake(err.Field()),
				Message: message,
			})
		}
		s.logger.Error("Validation error", zap.Any("validationErrors", validationErrors))
		return entity.AuthTokens{}, &serr.ServiceError{
			Message:    "Validation failed",
			StatusCode: http.StatusBadRequest,
			Details:    validationErrors,
		}
	}

	return s.domainService.Register(request.Email, request.Username, request.Password, userAgent)
}

func (s *UserApplicationService) Login(request *api.LoginRequest, userAgent string) (entity.AuthTokens, *serr.ServiceError) {

	if err := request.Validate(); err != nil {
		var validationErrors []api.ErrorContent
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("field '%s' is required", strcase.ToSnake(err.Field()))
			case "email":
				message = fmt.Sprintf("field '%s' must be a valid email address", strcase.ToSnake(err.Field()))
			case "username":
				message = fmt.Sprintf("field '%s' must be a valid username", strcase.ToSnake(err.Field()))
			default:
				message = fmt.Sprintf("field '%s' failed validation on the '%s' tag", strcase.ToSnake(err.Field()), err.Tag())
			}

			validationErrors = append(validationErrors, api.ErrorContent{
				Field:   strcase.ToSnake(err.Field()),
				Message: message,
			})
		}
		s.logger.Error("Validation error", zap.Any("validationErrors", validationErrors))
		return entity.AuthTokens{}, &serr.ServiceError{
			Message:    "Validation failed",
			StatusCode: http.StatusBadRequest,
			Details:    validationErrors,
		}
	}

	return s.domainService.Login(request.Email, request.Username, request.Password, userAgent)
}

func (s *UserApplicationService) Logout(refreshToken string) *serr.ServiceError {
	return s.domainService.Logout(refreshToken)
}
