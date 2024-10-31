package config

import (
	"github.com/Iamirup/whaler/backend/microservice/auth/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservice/auth/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservice/auth/pkg/token"
	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/infrastructure/repository"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	RDBMS      *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
	Token      *token.Config      `koanf:"token"`
}
