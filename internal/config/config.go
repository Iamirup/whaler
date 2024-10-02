package config

import (
	"github.com/Iamirup/whaler/internal/repository"
	"github.com/Iamirup/whaler/pkg/logger"
	"github.com/Iamirup/whaler/pkg/rdbms"
	"github.com/Iamirup/whaler/pkg/token"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	RDBMS      *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
	Token      *token.Config      `koanf:"token"`
}
