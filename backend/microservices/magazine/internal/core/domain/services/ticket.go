package services

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/magazine/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/microservices/magazine/pkg/errors"
	"go.uber.org/zap"
)

type NewsService struct {
	newsPersistencePort ports.NewsPersistencePort
	logger              *zap.Logger
	token               token.Token
}

func NewNewsService(
	newsPersistencePort ports.NewsPersistencePort,
	logger *zap.Logger, token token.Token) ports.NewsServicePort {

	return &NewsService{
		newsPersistencePort: newsPersistencePort,
		logger:              logger,
		token:               token,
	}
}

func (s *NewsService) AddNews(title, content string) (entity.UUID, *serr.ServiceError) {

	newsEntity := &entity.News{
		Title:   title,
		Content: content,
	}

	err := s.newsPersistencePort.CreateNews(newsEntity)
	if err != nil {
		s.logger.Error("Error happened while creating the news", zap.Error(err))
		return "", &serr.ServiceError{Message: "Error happened while creating the news", StatusCode: http.StatusInternalServerError}
	} else if newsEntity.NewsId == "" {
		s.logger.Error("Error invalid news id created", zap.Any("news", newsEntity))
		return "", &serr.ServiceError{Message: "Error invalid news id created", StatusCode: http.StatusInternalServerError}
	}

	return newsEntity.NewsId, nil
}

func (s *NewsService) MyNews(encryptedCursor string, limit int) ([]entity.News, string, *serr.ServiceError) {

	news, newEncryptedCursor, err := s.newsPersistencePort.GetNews(encryptedCursor, limit)
	if err != nil {
		s.logger.Error("Something went wrong in retrieving news", zap.Error(err))
		return nil, "", &serr.ServiceError{Message: "Something went wrong in retrieving news", StatusCode: http.StatusInternalServerError}
	}

	return news, newEncryptedCursor, nil
}
