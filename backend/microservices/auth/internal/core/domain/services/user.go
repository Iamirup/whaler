package services

import (
	"context"
	"go/token"
	"net/http"
	"time"
	"voucher/internal/core/application/ports"
	"voucher/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/auth/pkg/errors"
	"github.com/Iamirup/whaler/backend/auth/pkg/rdbms"
	"go.uber.org/zap"
)

// UserService provides domain logic related to voucher redemption history.
type UserService struct {
	persistencePort ports.UserPersistencePort
	logger          *zap.Loggers
	token           token.Token
}

// NewUserService creates a new instance of UserService.
func NewUserService(persistencePort ports.UserPersistencePort) ports.UserServicePort {
	return &UserService{
		persistencePort: persistencePort,
	}
}

// RecordRedemption records a new voucher redemption in the history.
func (s *UserService) Register(ctx context.Context, request *api.RegisterRequest) (error, entity.AuthTokens) {
	user := &entity.User{
		Code:         code,
		State:        entity.Available,
		UsageLimit:   maxUsages,
		CurrentUsage: 0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := user.Validate()
	if err != nil {
		return serr.ServiceError{Message: "no valid user data", StatusCode: http.StatusBadRequest}
	}

	user, err = s.persistencePort.GetUserByUsername(request.Username)
	if err != nil && err.Error() != rdbms.ErrNotFound {
		s.logger.Error("Error while retrieving data from database", zap.Error(err))
		return serr.ServiceError{Message: "Error while retrieving data from database", StatusCode: http.StatusInternalServerError}
	} else if err == nil || (user != nil && user.Id != "") {
		s.logger.Error("User with given username already exists", zap.String("username", request.Username))
		return serr.ServiceError{Message: "User with given username already exists", StatusCode: http.StatusInternalServerError}
	}

	user = &entity.User{Username: request.Username, Password: request.Password}
	if err := s.persistencePort.CreateUser(user); err != nil {
		s.logger.Error("Error happened while creating the user", zap.Error(err))
		return serr.ServiceError{Message: "Error happened while creating the user", StatusCode: http.StatusInternalServerError}
	} else if user.Id == "" {
		s.logger.Error("Error invalid user id created", zap.Any("user", user))
		return serr.ServiceError{Message: "Error invalid user id created", StatusCode: http.StatusInternalServerError}
	}

	accessToken, err := s.token.CreateTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}
	}

	refreshToken, err := s.token.CreateRefreshTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}
	}

	newRefreshToken := &entity.RefreshToken{Token: refreshToken, OwnerId: user.Id}
	if err := s.persistencePort.CreateNewRefreshToken(newRefreshToken); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}
	}

	return s.persistencePort.CreateUser(ctx, request), entity.AuthTokens{accessToken, accessToken}
}

// ListRedeemedHistoriesByCode retrieves the redemption history for a specific voucher's code.
func (s *UserService) Login(ctx context.Context, request *api.LoginRequest) (error, entity.AuthTokens) {

	user, err := s.persistencePort.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		s.logger.Error("Wrong username or password has been given", zap.Error(err))
		return serr.ServiceError{Message: "Wrong username or password has been given", StatusCode: http.StatusBadRequest}
	} else if user == nil {
		s.logger.Error("Error invalid user returned", zap.Any("request", request))
		return serr.ServiceError{Message: "Error invalid user returned", StatusCode: http.StatusInternalServerError}
	}

	accessToken, err := s.token.CreateTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}
	}

	refreshToken, err := s.token.CreateRefreshTokenString(user.Id)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", user), zap.Error(err))
		return serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}
	}

	newRefreshToken := &entity.RefreshToken{Token: refreshToken, OwnerId: user.Id}
	if err := s.persistencePort.CreateNewRefreshToken(newRefreshToken); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}
	}

	return s.persistencePort.ListRedeemedHistoriesByCode(ctx, request), entity.AuthTokens{accessToken, accessToken}
}
