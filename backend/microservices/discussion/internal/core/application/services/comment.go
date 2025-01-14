package services

import (
	"fmt"
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/discussion/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/discussion/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/domain/entity"
)

type CommentApplicationService struct {
	domainService ports.CommentServicePort
	logger        *zap.Logger
}

func NewCommentApplicationService(domainService ports.CommentServicePort, logger *zap.Logger) *CommentApplicationService {
	return &CommentApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *CommentApplicationService) NewComment(request *api.NewCommentRequest, username string) (int64, *serr.ServiceError) {

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
		return 0, &serr.ServiceError{
			Message:    "Validation failed",
			StatusCode: http.StatusBadRequest,
			Details:    validationErrors,
		}
	}

	return s.domainService.NewComment(request.Text, request.Currency, username)
}

func (s *CommentApplicationService) GetComments(currencyTopic, encryptedCursor string, limit int) ([]entity.Comment, string, *serr.ServiceError) {
	return s.domainService.GetComments(currencyTopic, encryptedCursor, limit)
}
