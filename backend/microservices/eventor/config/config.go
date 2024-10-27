package config

import (
	"github.com/Iamirup/whaler/backend/eventor/internal/repository"
	"github.com/Iamirup/whaler/backend/eventor/pkg/logger"
	"github.com/Iamirup/whaler/backend/eventor/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/eventor/pkg/token"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	RDBMS      *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
	Token      *token.Config      `koanf:"token"`
}
