package config

import (
	"github.com/Iamirup/whaler/backend/microservices/blog/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/token"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	RDBMS      *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
	Token      *token.Config      `koanf:"token"`
}
