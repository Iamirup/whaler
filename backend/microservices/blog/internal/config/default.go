package config

import (
	"time"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/token"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
		RDBMS: &rdbms.Config{
			Host:     "rdbms",
			Port:     5432,
			Username: "user",
			Password: "pass",
			Database: "db",
		},
		Repository: &repository.Config{
			CursorSecret: "A?D(G-KaPdSgVkYp",
			Limit: struct {
				Min int "koanf:\"min\""
				Max int "koanf:\"max\""
			}{12, 48},
			Pepper: "WqL2kwU3pQav1Al",
		},
		Token: &token.Config{
			PublicPem: "-----BEGIN PUBLIC KEY-----\n" +
				"MCowBQYDK2VwAyEAqQsZ5iRNP3kdpNn3V/db9o/WkYHY8kkwQqCZGcDvJ+g=\n" +
				"-----END PUBLIC KEY-----",
			AccessTokenExpiration: 10 * time.Minute,
		},
	}
}
