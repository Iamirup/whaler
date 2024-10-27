package config

import (
	"github.com/Iamirup/whaler/backend/auth/internal/repository"
	"github.com/Iamirup/whaler/backend/auth/pkg/logger"
	"github.com/Iamirup/whaler/backend/auth/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/auth/pkg/token"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	RDBMS      *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
	Token      *token.Config      `koanf:"token"`
}
