package services

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/microservices/discussion/pkg/errors"
	"go.uber.org/zap"
)

type CommentService struct {
	commentPersistencePort ports.CommentPersistencePort
	logger                 *zap.Logger
	token                  token.Token
}

func NewCommentService(
	commentPersistencePort ports.CommentPersistencePort,
	logger *zap.Logger, token token.Token) ports.CommentServicePort {

	return &CommentService{
		commentPersistencePort: commentPersistencePort,
		logger:                 logger,
		token:                  token,
	}
}

func (s *CommentService) NewComment(text, currency string, username string) (entity.UUID, *serr.ServiceError) {

	commentEntity := &entity.Comment{
		Username: username,
		Text:     text,
		Currency: currency,
	}

	err := s.commentPersistencePort.AddComment(commentEntity)
	if err != nil {
		s.logger.Error("Error happened while creating the comment", zap.Error(err))
		return "", &serr.ServiceError{Message: "Error happened while creating the comment", StatusCode: http.StatusInternalServerError}
	} else if commentEntity.CommentId == "" {
		s.logger.Error("Error invalid comment id created", zap.Any("comment", commentEntity))
		return "", &serr.ServiceError{Message: "Error invalid comment id created", StatusCode: http.StatusInternalServerError}
	}

	return commentEntity.CommentId, nil
}

func (s *CommentService) GetComments(currencyTopic, encryptedCursor string, limit int) ([]entity.Comment, string, *serr.ServiceError) {

	comments, newEncryptedCursor, err := s.commentPersistencePort.GetComments(currencyTopic, encryptedCursor, limit)
	if err != nil {
		s.logger.Error("Something went wrong in retrieving comments", zap.Error(err))
		return nil, "", &serr.ServiceError{Message: "Something went wrong in retrieving comments", StatusCode: http.StatusInternalServerError}
	}

	return comments, newEncryptedCursor, nil
}
