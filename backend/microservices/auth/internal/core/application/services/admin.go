package services

import (
	"fmt"
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"

	"github.com/iancoleman/strcase"
)

type AdminApplicationService struct {
	domainService ports.AdminServicePort
	logger        *zap.Logger
}

func NewAdminApplicationService(domainService ports.AdminServicePort, logger *zap.Logger) *AdminApplicationService {
	return &AdminApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *AdminApplicationService) AddAdmin(request *api.AddAdminRequest) *serr.ServiceError {
	if err := request.Validate(); err != nil {
		var validationErrors []api.ErrorContent
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("field '%s' is required", strcase.ToSnake(err.Field()))
			default:
				message = fmt.Sprintf("field '%s' failed validation on the '%s' tag", strcase.ToSnake(err.Field()), err.Tag())
			}
			validationErrors = append(validationErrors, api.ErrorContent{
				Field:   strcase.ToSnake(err.Field()),
				Message: message,
			})
		}
		s.logger.Error("Validation error", zap.Any("validationErrors", validationErrors))
		return &serr.ServiceError{
			Message:    "Validation failed",
			StatusCode: http.StatusBadRequest,
			Details:    validationErrors,
		}
	}
	return s.domainService.AddAdmin(request.UserId)
}

func (s *AdminApplicationService) DeleteAdmin(request *api.DeleteAdminRequest) *serr.ServiceError {
	if err := request.Validate(); err != nil {
		var validationErrors []api.ErrorContent
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("field '%s' is required", strcase.ToSnake(err.Field()))
			default:
				message = fmt.Sprintf("field '%s' failed validation on the '%s' tag", strcase.ToSnake(err.Field()), err.Tag())
			}
			validationErrors = append(validationErrors, api.ErrorContent{
				Field:   strcase.ToSnake(err.Field()),
				Message: message,
			})
		}
		s.logger.Error("Validation error", zap.Any("validationErrors", validationErrors))
		return &serr.ServiceError{
			Message:    "Validation failed",
			StatusCode: http.StatusBadRequest,
			Details:    validationErrors,
		}
	}
	return s.domainService.DeleteAdmin(request.AdminId)
}
