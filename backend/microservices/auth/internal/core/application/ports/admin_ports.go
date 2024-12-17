package ports

import (
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	serr "github.com/Iamirup/whaler/backend/microservices/auth/pkg/errors"
)

type (
	// AdminPersistencePort defines the methods for interacting with admin data
	AdminPersistencePort interface {
		// CreateNewAdmin adds a new admin to the database
		CreateNewAdmin(userId entity.UUID) error

		// DeleteAdmin removes an admin from database
		RemoveAdmin(adminId entity.UUID) error
	}

	// AdminServicePort defines the methods for interacting with admin services
	AdminServicePort interface {
		AddAdmin(userId entity.UUID) *serr.ServiceError
		DeleteAdmin(adminId entity.UUID) *serr.ServiceError
	}
)
