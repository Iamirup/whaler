package services

import (
	"context"
	"net/http"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"

	"github.com/Iamirup/whaler/backend/microservice/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservice/auth/internal/core/domain/entity"
	api "github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest/dto"

	serr "github.com/Iamirup/whaler/backend/microservice/auth/pkg/errors"
	"github.com/Iamirup/whaler/backend/microservice/auth/pkg/rdbms"
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
func (s *UserService) Register(ctx context.Context, request *api.RegisterRequest) (error, entity.AuthTokens) {
	user := &entity.User{
		Username:  request.Username,
		Password:  request.Password,
		CreatedAt: time.Now(),
	}
	err := user.Validate()
	if err != nil {
		return serr.ServiceError{Message: "no valid user data", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	}

	user, err = s.userPersistencePort.GetUserByUsername(request.Username)
	if err != nil && err.Error() != rdbms.ErrNotFound {
		s.logger.Error("Error while retrieving data from database", zap.Error(err))
		return serr.ServiceError{Message: "Error while retrieving data from database", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	} else if err == nil || (user != nil && user.Id != "") {
		s.logger.Error("User with given username already exists", zap.String("username", request.Username))
		return serr.ServiceError{Message: "User with given username already exists", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	user = &entity.User{Username: request.Username, Password: request.Password}
	if err := s.userPersistencePort.CreateUser(user); err != nil {
		s.logger.Error("Error happened while creating the user", zap.Error(err))
		return serr.ServiceError{Message: "Error happened while creating the user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	} else if user.Id == "" {
		s.logger.Error("Error invalid user id created", zap.Any("user", user))
		return serr.ServiceError{Message: "Error invalid user id created", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	accessToken, err := s.token.CreateTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	refreshToken, err := s.token.CreateRefreshTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	newRefreshToken := &entity.RefreshToken{Token: refreshToken, OwnerId: user.Id}
	if err := s.refreshTokenPersistencePort.CreateNewRefreshToken(newRefreshToken); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	return s.userPersistencePort.CreateUser(ctx, request), entity.AuthTokens{AccessToken: accessToken, RefreshToken: refreshToken}
}

// ListRedeemedHistoriesByCode retrieves the redemption history for a specific voucher's code.
func (s *UserService) Login(ctx context.Context, request *api.LoginRequest) (error, entity.AuthTokens) {

	user, err := s.userPersistencePort.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		s.logger.Error("Wrong username or password has been given", zap.Error(err))
		return serr.ServiceError{Message: "Wrong username or password has been given", StatusCode: http.StatusBadRequest}, entity.AuthTokens{}
	} else if user == nil {
		s.logger.Error("Error invalid user returned", zap.Any("request", request))
		return serr.ServiceError{Message: "Error invalid user returned", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	accessToken, err := s.token.CreateTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	refreshToken, err := s.token.CreateRefreshTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	newRefreshToken := &entity.RefreshToken{Token: refreshToken, OwnerId: user.Id}
	if err := s.refreshTokenPersistencePort.CreateNewRefreshToken(newRefreshToken); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}, entity.AuthTokens{}
	}

	return nil, entity.AuthTokens{AccessToken: accessToken, RefreshToken: refreshToken}
}
