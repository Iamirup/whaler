package services

import (
	"fmt"
	"net/http"
	"net/url"

	serr "github.com/Iamirup/whaler/backend/microservices/blog/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/blog/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"
)

type ArticleApplicationService struct {
	domainService ports.ArticleServicePort
	logger        *zap.Logger
}

func NewArticleApplicationService(domainService ports.ArticleServicePort, logger *zap.Logger) *ArticleApplicationService {
	return &ArticleApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *ArticleApplicationService) GetAnArticle(urlPath string) (*entity.Article, *serr.ServiceError) {
	return s.domainService.GetAnArticle(urlPath)
}

func (s *ArticleApplicationService) GetAllArticles(encryptedCursor string, limit int) ([]entity.Article, string, *serr.ServiceError) {
	return s.domainService.GetAllArticles(encryptedCursor, limit)
}

func (s *ArticleApplicationService) GetMyArticles(encryptedCursor string, limit int, authorId entity.UUID) ([]entity.Article, string, *serr.ServiceError) {
	return s.domainService.GetMyArticles(encryptedCursor, limit, authorId)
}

func (s *ArticleApplicationService) NewArticle(request *api.NewArticleRequest, userId entity.UUID, username string) (entity.UUID, *serr.ServiceError) {

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

	encodedURL := url.PathEscape(request.Title)

	return s.domainService.NewArticle(request.Title, request.Content, encodedURL, userId, username)
}

func (s *ArticleApplicationService) UpdateArticle(request *api.UpdateArticleRequest, userId entity.UUID) *serr.ServiceError {
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

	encodedURL := url.PathEscape(request.Title)

	return s.domainService.UpdateArticle(request.ArticleId, request.Title, encodedURL, request.Content, userId)
}

func (s *ArticleApplicationService) DeleteArticle(request *api.DeleteArticleRequest, userId entity.UUID) *serr.ServiceError {
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
	return s.domainService.DeleteArticle(request.ArticleId, userId)
}
