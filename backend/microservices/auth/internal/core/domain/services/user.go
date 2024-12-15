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

func (s *UserService) Register(email, username, password string) (entity.AuthTokens, *serr.ServiceError) {

	userEntity := &entity.User{
		Email:    email,
		Username: username,
		Password: password,
	}

	if err := s.userPersistencePort.CreateUser(userEntity); err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			s.logger.Error("User with given email already exists", zap.String("email", email))
			return entity.AuthTokens{}, &serr.ServiceError{Message: "User with given email already exists", StatusCode: http.StatusBadRequest}
		} else if strings.Contains(err.Error(), "users_username_key") {
			s.logger.Error("User with given username already exists", zap.String("username", username))
			return entity.AuthTokens{}, &serr.ServiceError{Message: "User with given username already exists", StatusCode: http.StatusBadRequest}
		}
		s.logger.Error("Error happened while creating the user", zap.Error(err))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error happened while creating the user", StatusCode: http.StatusInternalServerError}
	} else if userEntity.Id == "" {
		s.logger.Error("Error invalid user id created", zap.Any("user", userEntity))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error invalid user id created", StatusCode: http.StatusInternalServerError}
	}

	isAdmin, err := s.refreshTokenPersistencePort.CheckIfIsAdmin(userEntity.Id)
	if err != nil {
		s.logger.Error("something went wrong")
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Something went wrong! please try again later", StatusCode: http.StatusInternalServerError}
	}

	newAccessToken, err := s.token.CreateTokenString(userEntity.Id, userEntity.Username, isAdmin)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", userEntity), zap.Error(err))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}
	}

	newRefreshToken, err := s.token.CreateRefreshTokenString(userEntity.Id)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", userEntity), zap.Error(err))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}
	}

	refreshTokenEntity := &entity.RefreshToken{Token: newRefreshToken, OwnerId: userEntity.Id}
	if err := s.refreshTokenPersistencePort.CreateNewRefreshToken(refreshTokenEntity); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}
	}

	return entity.AuthTokens{AccessToken: newAccessToken, RefreshToken: newRefreshToken}, nil
}

func (s *UserService) Login(email, username, password, possibleRefreshToken string) (entity.AuthTokens, *serr.ServiceError) {

	var user *entity.User
	var err error

	if possibleRefreshToken != "" {
		userId := s.refreshTokenPersistencePort.CheckRefreshTokenExistsInDB(possibleRefreshToken)
		if userId != "" {
			s.logger.Error("This user is already logged in", zap.String("email", email))
			return entity.AuthTokens{}, &serr.ServiceError{Message: "You already logged in", StatusCode: http.StatusBadRequest}
		}
	}

	if strings.TrimSpace(email) != "" {
		user, err = s.userPersistencePort.GetUserByEmailAndPassword(email, password)
		if err != nil {
			s.logger.Error("Wrong email or password has been given", zap.Error(err))
			return entity.AuthTokens{}, &serr.ServiceError{Message: "Wrong email or password has been given", StatusCode: http.StatusBadRequest}
		} else if user == nil {
			s.logger.Error("Error invalid user returned", zap.Any("request", fmt.Sprintf("%s - %s", email, password)))
			return entity.AuthTokens{}, &serr.ServiceError{Message: "Error invalid user returned", StatusCode: http.StatusInternalServerError}
		}
	} else {
		user, err = s.userPersistencePort.GetUserByUsernameAndPassword(username, password)
		if err != nil {
			s.logger.Error("Wrong username or password has been given", zap.Error(err))
			return entity.AuthTokens{}, &serr.ServiceError{Message: "Wrong username or password has been given", StatusCode: http.StatusBadRequest}
		} else if user == nil {
			s.logger.Error("Error invalid user returned", zap.Any("request", fmt.Sprintf("%s - %s", username, password)))
			return entity.AuthTokens{}, &serr.ServiceError{Message: "Error invalid user returned", StatusCode: http.StatusInternalServerError}
		}
	}

	isAdmin, err := s.refreshTokenPersistencePort.CheckIfIsAdmin(user.Id)
	if err != nil {
		s.logger.Error("something went wrong")
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Something went wrong! please try again later", StatusCode: http.StatusInternalServerError}
	}

	newAccessToken, err := s.token.CreateTokenString(user.Id, user.Username, isAdmin)
	if err != nil {
		s.logger.Error("Error creating JWT access token for user", zap.Any("user", user), zap.Error(err))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error creating JWT access token for user", StatusCode: http.StatusInternalServerError}
	}

	newRefreshToken, err := s.token.CreateRefreshTokenString(user.Id)
	fmt.Println("newRefreshToken: ", newRefreshToken)
	if err != nil {
		s.logger.Error("Error creating JWT refresh token for user", zap.Any("user", user), zap.Error(err))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error creating JWT refresh token for user", StatusCode: http.StatusInternalServerError}
	}

	refreshTokenEntity := &entity.RefreshToken{Token: newRefreshToken, OwnerId: user.Id}
	if err := s.refreshTokenPersistencePort.CreateNewRefreshToken(refreshTokenEntity); err != nil {
		s.logger.Error("Error happened while adding the refresh token", zap.Error(err))
		return entity.AuthTokens{}, &serr.ServiceError{Message: "Error happened while adding the refresh token", StatusCode: http.StatusInternalServerError}
	}

	return entity.AuthTokens{AccessToken: newAccessToken, RefreshToken: newRefreshToken}, nil
}

func (s *UserService) Logout(refreshToken string) *serr.ServiceError {

	err := s.refreshTokenPersistencePort.RemoveRefreshToken(refreshToken)
	if err != nil {
		s.logger.Error("Error invalid refresh token", zap.Error(err))
		return &serr.ServiceError{Message: "Invalid refresh token, abnormal activity was detected. please login again", StatusCode: http.StatusInternalServerError}
	}

	return nil
}
