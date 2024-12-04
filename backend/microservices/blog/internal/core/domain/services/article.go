package services

import (
	"net/http"
	"strings"

	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/microservices/blog/pkg/errors"
	"go.uber.org/zap"
)

type ArticleService struct {
	articlePersistencePort ports.ArticlePersistencePort
	logger                 *zap.Logger
	token                  token.Token
}

func NewArticleService(
	articlePersistencePort ports.ArticlePersistencePort,
	logger *zap.Logger, token token.Token) ports.ArticleServicePort {
	return &ArticleService{
		articlePersistencePort: articlePersistencePort,
		logger:                 logger,
		token:                  token,
	}
}

func (s *ArticleService) GetAnArticle(urlPath string) (*entity.Article, *serr.ServiceError) {

	article, err := s.articlePersistencePort.GetAnArticle(urlPath)
	if err != nil {
		s.logger.Error("There is no article with this url path", zap.Error(err))
		return &entity.Article{}, &serr.ServiceError{Message: "There is no article with this url path", StatusCode: http.StatusNotFound}
	}

	return article, nil
}

func (s *ArticleService) GetArticles(encryptedCursor string, limit int) ([]entity.Article, string, *serr.ServiceError) {

	articles, newEncryptedCursor, err := s.articlePersistencePort.GetArticles(encryptedCursor, limit)
	if err != nil {
		s.logger.Error("Something went wrong in retrieving articles", zap.Error(err))
		return nil, "", &serr.ServiceError{Message: "Something went wrong in retrieving articles", StatusCode: http.StatusInternalServerError}
	}

	return articles, newEncryptedCursor, nil
}

func (s *ArticleService) NewArticle(title, content, urlPath string, authorId entity.UUID, authorUsername string) (entity.UUID, *serr.ServiceError) {

	articleEntity := &entity.Article{
		UrlPath:        urlPath,
		Title:          title,
		Content:        content,
		AuthorId:       authorId,
		AuthorUsername: authorUsername,
	}

	err := s.articlePersistencePort.CreateArticle(articleEntity)
	if err != nil {
		s.logger.Error("Error happened while creating the user", zap.Error(err))
		return "", &serr.ServiceError{Message: "Error happened while creating the user", StatusCode: http.StatusInternalServerError}
	} else if articleEntity.ArticleId == "" {
		s.logger.Error("Error invalid article id created", zap.Any("article", articleEntity))
		return "", &serr.ServiceError{Message: "Error invalid article id created", StatusCode: http.StatusInternalServerError}
	}

	return articleEntity.ArticleId, nil
}

func (s *ArticleService) UpdateArticle(articleId entity.UUID, title, content string, authorId entity.UUID) *serr.ServiceError {

	err := s.articlePersistencePort.CheckIfIsAuthorById(articleId, authorId)
	if err != nil {
		if err.Error() == "you are not the author of this article" {
			s.logger.Error("You are not the author of this article!", zap.Error(err))
			return &serr.ServiceError{Message: "You are not the author of this article!", StatusCode: http.StatusForbidden}
		}
		s.logger.Error("an error occured", zap.Error(err))
		return &serr.ServiceError{Message: "an error occured", StatusCode: http.StatusInternalServerError}
	}

	if strings.TrimSpace(title) != "" {
		if err := s.articlePersistencePort.UpdateArticleTitle(articleId, title); err != nil {
			s.logger.Error("Wrong article has been given", zap.Error(err))
			return &serr.ServiceError{Message: "Wrong article has been given", StatusCode: http.StatusBadRequest}
		}
	}
	if strings.TrimSpace(content) != "" {
		if err := s.articlePersistencePort.UpdateArticleContent(articleId, content); err != nil {
			s.logger.Error("Wrong article has been given", zap.Error(err))
			return &serr.ServiceError{Message: "Wrong article has been given", StatusCode: http.StatusBadRequest}
		}
	}

	return nil
}

func (s *ArticleService) DeleteArticle(articleId entity.UUID, authorId entity.UUID) *serr.ServiceError {

	err := s.articlePersistencePort.CheckIfIsAuthorById(articleId, authorId)
	if err != nil {
		if err.Error() == "you are not the author of this article" {
			s.logger.Error("You are not the author of this article!", zap.Error(err))
			return &serr.ServiceError{Message: "You are not the author of this article!", StatusCode: http.StatusForbidden}
		}
		s.logger.Error("an error occured", zap.Error(err))
		return &serr.ServiceError{Message: "an error occured", StatusCode: http.StatusInternalServerError}
	}

	if err := s.articlePersistencePort.RemoveArticle(articleId); err != nil {
		s.logger.Error("Error creating a reply to the article", zap.Error(err))
		return &serr.ServiceError{Message: "Error creating a reply to the article", StatusCode: http.StatusInternalServerError}
	}

	return nil
}
