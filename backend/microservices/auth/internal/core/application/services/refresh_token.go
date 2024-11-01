package services

import "github.com/Iamirup/whaler/backend/microservices/auth/internal/core/application/ports"

type RefreshTokenApplicationService struct {
	domainService ports.RefreshTokenServicePort
}
