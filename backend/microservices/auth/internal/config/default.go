package config

import (
	"time"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/infrastructure/repository"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/logger"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/token"
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
			PrivatePem: "-----BEGIN PRIVATE KEY-----\n" +
				"MC4CAQAwBQYDK2VwBCIEINyMNS8h9M9HO73Tg1BPr53p//qlqylO+wPKN8GrlsX7\n" +
				"-----END PRIVATE KEY-----",
			PublicPem: "-----BEGIN PUBLIC KEY-----\n" +
				"MCowBQYDK2VwAyEAqQsZ5iRNP3kdpNn3V/db9o/WkYHY8kkwQqCZGcDvJ+g=\n" +
				"-----END PUBLIC KEY-----",
			AccessTokenExpiration:  10 * time.Minute,
			RefreshTokenExpiration: 24 * 14 * time.Hour,
		},
	}
}
