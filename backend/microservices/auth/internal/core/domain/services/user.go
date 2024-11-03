package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/rdbms"
	"go.uber.org/zap"
)

// UserService provides domain logic related to voucher redemption history.
type UserService struct {
	userPersistencePort         ports.UserPersistencePort
	refreshTokenPersistencePort ports.RefreshTokenPersistencePort
	logger                      *zap.Logger
	token                       token.Token
}

// NewUserService creates a new instance of UserService.
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

// RecordRedemption records a new voucher redemption in the history.
func (s *UserService) Register(email, username, password string) (*serr.ServiceError, entity.AuthTokens) {
	user := &entity.User{
		Email:     email,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}
	err := user.Validate()
	if err != nil {
		return &serr.ServiceError{Message: "no valid user data", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	user, err = s.userPersistencePort.GetUserByUsername(username)
	if err != nil && err.Error() != rdbms.ErrNotFound {
		s.logger.Error("Error while retrieving data from database", zap.Error(err))
		return &serr.ServiceError{Message: "Error while retrieving data from database", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	} else if err == nil || (user != nil && user.Id != "") {
		s.logger.Error("User with given username already exists", zap.String("username", username))
		return &serr.ServiceError{Message: "User with given username already exists", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	user = &entity.User{Username: username, Password: password}
	if err := s.userPersistencePort.CreateUser(user); err != nil {
		s.logger.Error("Error happened while creating the user", zap.Error(err))
		return &serr.ServiceError{Message: "Error happened while creating the user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	} else if user.Id == "" {
		s.logger.Error("Error invalid user id created", zap.Any("user", user))
		return &serr.ServiceError{Message: "Error invalid user id created", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	accessToken, err := s.token.CreateTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return &serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	refreshToken, err := s.token.CreateRefreshTokenString(user.Id)
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

// ListRedeemedHistoriesByCode retrieves the redemption history for a specific voucher's code.
func (s *UserService) Login(username, password string) (*serr.ServiceError, entity.AuthTokens) {

	user, err := s.userPersistencePort.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		s.logger.Error("Wrong username or password has been given", zap.Error(err))
		return &serr.ServiceError{Message: "Wrong username or password has been given", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	} else if user == nil {
		s.logger.Error("Error invalid user returned", zap.Any("request", fmt.Sprintf("%s - %s", username, password)))
		return &serr.ServiceError{Message: "Error invalid user returned", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	accessToken, err := s.token.CreateTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return &serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	refreshToken, err := s.token.CreateRefreshTokenString(user.Id)
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
