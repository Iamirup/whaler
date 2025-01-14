package config

import (
	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/token"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	RDBMS      *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
	Token      *token.Config      `koanf:"token"`
}
