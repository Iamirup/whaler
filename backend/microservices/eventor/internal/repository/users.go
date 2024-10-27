package repository

import (
	"errors"

	"github.com/Iamirup/whaler/backend/eventor/internal/models"
	"github.com/Iamirup/whaler/backend/eventor/pkg/rdbms"
	"go.uber.org/zap"
)

const QueryCreateUser = `
INSERT INTO users(username, password) VALUES($1, $2)
RETURNING id;`

func (r *repository) CreateUser(user *models.User) error {

	if len(user.Username) == 0 || len(user.Password) == 0 {
		return errors.New("insufficient information for user")
	}

	hashedPassword, err := user.HashPassword()
	if err != nil {
		r.logger.Error("Error hashing password", zap.Error(err))
		return err
	}

	in := []any{user.Username, hashedPassword}
	out := []any{&user.Id}
	if err := r.rdbms.QueryRow(QueryCreateUser, in, out); err != nil {
		r.logger.Error("Error inserting author", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetUserDetailsById = `
SELECT *
FROM users
JOIN user_configs ON users.id = user_configs.user_id
WHERE
	users.id = $1 AND
	user_configs.id > $2 AND
	user_configs.name LIKE '%' || $3 || '%'
ORDER BY user_configs.id
FETCH NEXT $4 ROWS ONLY;`

const QueryGetUserByUsername = `
SELECT id, password, created_at
FROM users
WHERE username=$1;`

func (r *repository) GetUserByUsername(username string) (*models.User, error) {

	user := &models.User{Username: username}

	in := []interface{}{username}
	out := []interface{}{&user.Id, &user.Password, &user.CreatedAt}
	if err := r.rdbms.QueryRow(QueryGetUserByUsername, in, out); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			return nil, err
		}

		r.logger.Error("Error find user by username", zap.Error(err))
		return nil, err
	}

	return user, nil
}

const QueryGetUserByUsernameAndPassword = `
SELECT id, password, created_at 
FROM users
WHERE username=$1`

func (r *repository) GetUserByUsernameAndPassword(username, password string) (*models.User, error) {

	user := &models.User{Username: username}

	in := []interface{}{username}
	out := []interface{}{&user.Id, &user.Password, &user.CreatedAt}
	if err := r.rdbms.QueryRow(QueryGetUserByUsernameAndPassword, in, out); err != nil {
		r.logger.Error("Error finding user by username", zap.Error(err))
		return nil, err
	}

	if !user.CheckPasswordHash(password) {
		r.logger.Error("invalid username or password")
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
