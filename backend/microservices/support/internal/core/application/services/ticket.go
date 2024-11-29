package services

import (
	"fmt"
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/support/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/support/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"
)

type TicketApplicationService struct {
	domainService ports.TicketServicePort
	logger        *zap.Logger
}

func NewTicketApplicationService(domainService ports.TicketServicePort, logger *zap.Logger) *TicketApplicationService {
	return &TicketApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *TicketApplicationService) NewTicket(request *api.NewTicketRequest, userId entity.UUID, username string) (entity.UUID, *serr.ServiceError) {

	if err := request.Validate(); err != nil {
		var validationErrors []api.ErrorContent
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("field '%s' is required", strcase.ToSnake(err.Field()))
			case "min":
				message = fmt.Sprintf("field '%s' must be at least %s characters", strcase.ToSnake(err.Field()), strcase.ToSnake(err.Param()))
			case "max":
				message = fmt.Sprintf("field '%s' must be at most %s characters", strcase.ToSnake(err.Field()), strcase.ToSnake(err.Param()))
			default:
				message = fmt.Sprintf("field '%s' failed validation on the '%s' tag", strcase.ToSnake(err.Field()), err.Tag())
			}
			validationErrors = append(validationErrors, api.ErrorContent{
				Field:   strcase.ToSnake(err.Field()),
				Message: message,
			})
		}
		s.logger.Error("Validation error", zap.Any("validationErrors", validationErrors))
		return "", &serr.ServiceError{
			Message:    "Validation failed",
			StatusCode: http.StatusBadRequest,
			Details:    validationErrors,
		}
	}

	return s.domainService.NewTicket(request.Title, request.Content, userId, username)
}

func (s *TicketApplicationService) MyTickets(userId entity.UUID, encryptedCursor string, limit int) ([]entity.Ticket, string, *serr.ServiceError) {
	return s.domainService.MyTickets(userId, encryptedCursor, limit)
}

func (s *TicketApplicationService) ReplyToTicket(request *api.ReplyToTicketRequest) *serr.ServiceError {
	if err := request.Validate(); err != nil {
		var validationErrors []api.ErrorContent
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("field '%s' is required", strcase.ToSnake(err.Field()))
			case "min":
				message = fmt.Sprintf("field '%s' must be at least %s characters", strcase.ToSnake(err.Field()), strcase.ToSnake(err.Param()))
			case "max":
				message = fmt.Sprintf("field '%s' must be at most %s characters", strcase.ToSnake(err.Field()), strcase.ToSnake(err.Param()))
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

	return s.domainService.ReplyToTicket(request.TicketId, request.ReplyText)
}

func (s *TicketApplicationService) AllTicket(encryptedCursor string, limit int) ([]entity.Ticket, string, *serr.ServiceError) {
	return s.domainService.AllTicket(encryptedCursor, limit)
}
