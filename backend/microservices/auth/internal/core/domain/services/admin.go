package services

import (
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"

	"go.uber.org/zap"
)

type AdminService struct {
	adminPersistencePort ports.AdminPersistencePort
	logger               *zap.Logger
	token                token.Token
}

func NewAdminService(
	adminPersistencePort ports.AdminPersistencePort,
	logger *zap.Logger, token token.Token) ports.AdminServicePort {

	return &AdminService{
		adminPersistencePort: adminPersistencePort,
		logger:               logger,
		token:                token,
	}
}

func (s *AdminService) AddAdmin(userId entity.UUID) *serr.ServiceError {

	err := s.adminPersistencePort.CreateNewAdmin(userId)
	if err != nil {
		s.logger.Error("Error cant add new admin", zap.Error(err))
		return &serr.ServiceError{Message: "Can't add new admin", StatusCode: http.StatusInternalServerError}
	}

	return nil
}

func (s *AdminService) DeleteAdmin(adminId entity.UUID) *serr.ServiceError {

	err := s.adminPersistencePort.RemoveAdmin(adminId)
	if err != nil {
		s.logger.Error("Error cant add new admin", zap.Error(err))
		return &serr.ServiceError{Message: "Can't add new admin", StatusCode: http.StatusInternalServerError}
	}

	return nil
}
