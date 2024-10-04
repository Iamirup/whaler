package repository

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"

	"github.com/Iamirup/whaler/internal/models"
	"github.com/Iamirup/whaler/pkg/rdbms"
	"github.com/Iamirup/whaler/pkg/utils"
	"go.uber.org/zap"
)

type Repository interface {
	Migrate(models.Migrate) error

	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByUsernameAndPassword(username, password string) (*models.User, error)

	GetConfigById(userId string, configId uint64) (*models.UserConfig, error)
	UpdateConfig(userId string, config *models.UserConfig) error

	CreateNewRefreshToken(refresh_token *models.RefreshToken) error
	GetRefreshTokenById(ownerId string) (*models.RefreshToken, error)
}

type repository struct {
	logger *zap.Logger
	config *Config
	rdbms  rdbms.RDBMS
}

func New(logger *zap.Logger, cfg *Config, rdbms rdbms.RDBMS) Repository {
	return &repository{logger: logger, config: cfg, rdbms: rdbms}
}

// go:embed migrations
var migrations embed.FS

func (r *repository) Migrate(direction models.Migrate) error {
	fmt.Println("here: ", migrations)

	files, err := fs.ReadDir(migrations, ".")
	if err != nil {
		return fmt.Errorf("error reading migrations directory:\n%v", err)
	}

	fmt.Println("1: ", files)

	result := make([]string, 0, len(files)/2)

	for _, file := range files {
		splits := strings.Split(file.Name(), ".")
		if splits[1] == string(direction) {
			result = append(result, file.Name())
		}
	}

	fmt.Println("2: ", result)

	result = utils.Sort(result)

	fmt.Println("3: ", result)

	for index := 0; index < len(result); index++ {
		file := "migrations/"

		if direction == models.Up {
			file += result[index]
		} else {
			file += result[len(result)-index-1]
		}

		fmt.Println("4: ", file)

		data, err := fs.ReadFile(migrations, file)
		if err != nil {
			return fmt.Errorf("error reading migration file: %s\n%v", file, err)
		}

		fmt.Println("5: ", data)

		if err := r.rdbms.Execute(string(data), []any{}); err != nil {
			return fmt.Errorf("error migrating the file: %s\n%v", file, err)
		}
	}

	return nil
}
