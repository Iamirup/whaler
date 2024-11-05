package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"go.uber.org/zap"
)

type UserService struct {
	userPersistencePort         ports.UserPersistencePort
	refreshTokenPersistencePort ports.RefreshTokenPersistencePort
	logger                      *zap.Logger
	token                       token.Token
}

func NewUserService(
	userPersistencePort ports.UserPersistencePort,
	refreshTokenPersistencePort ports.RefreshTokenPersistencePort,
	logger *zap.Logger, token token.Token) ports.UserServicePort {

	return &UserService{
		userPersistencePort:         userPersistencePort,
		refreshTokenPersistencePort: refreshTokenPersistencePort,
		logger:                      logger,
		token:                       token,
	}
}

func (s *UserService) Register(email, username, password, userAgent string) (*serr.ServiceError, entity.AuthTokens) {

	userEntity := &entity.User{
		Email:    email,
		Username: username,
		Password: password,
	}

	if err := s.userPersistencePort.CreateUser(userEntity); err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			s.logger.Error("User with given email already exists", zap.String("email", email))
			return &serr.ServiceError{Message: "User with given email already exists", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
		} else if strings.Contains(err.Error(), "users_username_key") {
			s.logger.Error("User with given username already exists", zap.String("username", username))
			return &serr.ServiceError{Message: "User with given username already exists", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
		}
		s.logger.Error("Error happened while creating the user", zap.Error(err))
		return &serr.ServiceError{Message: "Error happened while creating the user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	} else if userEntity.Id == "" {
		s.logger.Error("Error invalid user id created", zap.Any("user", userEntity))
		return &serr.ServiceError{Message: "Error invalid user id created", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	accessToken, err := s.token.CreateTokenString(userEntity.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", userEntity), zap.Error(err))
		return &serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	newRefreshToken, err := s.token.CreateRefreshTokenString(userEntity.Id, userAgent)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", userEntity), zap.Error(err))
		return &serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	refreshTokenEntity := &entity.RefreshToken{Token: newRefreshToken, OwnerId: userEntity.Id}
	if err := s.refreshTokenPersistencePort.CreateNewRefreshToken(refreshTokenEntity); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return &serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	return nil, entity.AuthTokens{AccessToken: accessToken, RefreshToken: newRefreshToken}
}

func (s *UserService) Login(email, username, password, userAgent string) (*serr.ServiceError, entity.AuthTokens) {

	var user *entity.User
	var err error

	if strings.TrimSpace(email) != "" {
		user, err = s.userPersistencePort.GetUserByEmailAndPassword(email, password)
		if err != nil {
			s.logger.Error("Wrong email or password has been given", zap.Error(err))
			return &serr.ServiceError{Message: "Wrong email or password has been given", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
		} else if user == nil {
			s.logger.Error("Error invalid user returned", zap.Any("request", fmt.Sprintf("%s - %s", email, password)))
			return &serr.ServiceError{Message: "Error invalid user returned", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
		}
	} else {
		user, err = s.userPersistencePort.GetUserByUsernameAndPassword(username, password)
		if err != nil {
			s.logger.Error("Wrong username or password has been given", zap.Error(err))
			return &serr.ServiceError{Message: "Wrong username or password has been given", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
		} else if user == nil {
			s.logger.Error("Error invalid user returned", zap.Any("request", fmt.Sprintf("%s - %s", username, password)))
			return &serr.ServiceError{Message: "Error invalid user returned", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
		}
	}

	accessToken, err := s.token.CreateTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return &serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	refreshToken, err := s.token.CreateRefreshTokenString(user.Id, userAgent)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", user), zap.Error(err))
		return &serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	newRefreshToken := &entity.RefreshToken{Token: refreshToken, OwnerId: user.Id}
	if err := s.refreshTokenPersistencePort.CreateNewRefreshToken(newRefreshToken); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return &serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	return nil, entity.AuthTokens{AccessToken: accessToken, RefreshToken: refreshToken}
}

func (s *UserService) Logout(refreshToken string) *serr.ServiceError {
	err := s.refreshTokenPersistencePort.RemoveRefreshToken(refreshToken)
	if err != nil {
		s.logger.Error("Error invalid refresh token", zap.Error(err))
		return &serr.ServiceError{Message: "invalid refresh token, please login again", StatusCode: http.StatusInternalServerError}
	}

	return nil
}
