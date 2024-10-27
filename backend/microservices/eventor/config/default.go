package config

import (
	"time"

	"github.com/Iamirup/whaler/backend/eventor/internal/repository"
	"github.com/Iamirup/whaler/backend/eventor/pkg/logger"
	"github.com/Iamirup/whaler/backend/eventor/pkg/rdbms"
	"github.com/Iamirup/whaler/backend/eventor/pkg/token"
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
			// CursorSecret: "A?D(G-KaPdSgVkYp",
			// Limit: struct {
			// 	Min int "koanf:\"min\""
			// 	Max int "koanf:\"max\""
			// }{12, 48},
		},
		Token: &token.Config{
			PrivatePem: "-----BEGIN PRIVATE KEY-----\n" +
				"MC4CAQAwBQYDK2VwBCIEINyMNS8h9M9HO73Tg1BPr53p//qlqylO+wPKN8GrlsX7\n" +
				"-----END PRIVATE KEY-----",
			PublicPem: "-----BEGIN PUBLIC KEY-----\n" +
				"MCowBQYDK2VwAyEAqQsZ5iRNP3kdpNn3V/db9o/WkYHY8kkwQqCZGcDvJ+g=\n" +
				"-----END PUBLIC KEY-----",
			AccessTokenExpiration:  10 * time.Minute,
			RefreshTokenExpiration: 24 * 7 * time.Hour,
		},
	}
}
