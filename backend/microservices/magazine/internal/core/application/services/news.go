package services

import (
	"fmt"
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/magazine/pkg/errors"
	"github.com/go-playground/validator"
	"github.com/iancoleman/strcase"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/magazine/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/domain/entity"
)

type NewsApplicationService struct {
	domainService ports.NewsServicePort
	logger        *zap.Logger
}

func NewNewsApplicationService(domainService ports.NewsServicePort, logger *zap.Logger) *NewsApplicationService {
	return &NewsApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *NewsApplicationService) AddNews(request *api.AddNewsRequest) (entity.UUID, *serr.ServiceError) {

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

	return s.domainService.AddNews(request.Title, request.Content)
}

func (s *NewsApplicationService) SeeNews(encryptedCursor string, limit int) ([]entity.News, string, *serr.ServiceError) {
	return s.domainService.SeeNews(encryptedCursor, limit)
}
