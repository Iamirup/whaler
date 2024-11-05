package repository

import (
	"errors"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/auth/pkg/rdbms"
	"go.uber.org/zap"
)

const QueryCreateUser = `
INSERT INTO users(email, username, password) VALUES($1, $2, $3)
RETURNING id;`

func (r *userRepository) CreateUser(user *entity.User) error {

	if len(user.Email) == 0 || len(user.Username) == 0 || len(user.Password) == 0 {
		return errors.New("insufficient information for user")
	}

	hashedPassword, err := user.HashPassword()
	if err != nil {
		r.logger.Error("Error hashing password", zap.Error(err))
		return err
	}

	in := []any{user.Email, user.Username, hashedPassword}
	out := []any{&user.Id}
	if err := r.rdbms.QueryRow(QueryCreateUser, in, out); err != nil {
		r.logger.Error("Error inserting user", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetUserByUsername = `
SELECT id
FROM users
WHERE username=$1;`

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {

	user := &entity.User{Username: username}

	in := []interface{}{username}
	out := []interface{}{&user.Id}
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
SELECT id
FROM users
WHERE username=$1`

func (r *userRepository) GetUserByUsernameAndPassword(username, password string) (*entity.User, error) {

	user := &entity.User{Username: username}

	in := []interface{}{username}
	out := []interface{}{&user.Id}
	if err := r.rdbms.QueryRow(QueryGetUserByUsernameAndPassword, in, out); err != nil {
		r.logger.Error("Error finding user by username and password", zap.Error(err))
		return nil, err
	}

	if !user.CheckPasswordHash(password) {
		r.logger.Error("invalid username or password")
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

const QueryGetUserByEmailAndPassword = `
SELECT id
FROM users
WHERE email=$1`

func (r *userRepository) GetUserByEmailAndPassword(email, password string) (*entity.User, error) {

	user := &entity.User{Email: email}

	in := []interface{}{email}
	out := []interface{}{&user.Id}
	if err := r.rdbms.QueryRow(QueryGetUserByEmailAndPassword, in, out); err != nil {
		r.logger.Error("Error finding user by email and password", zap.Error(err))
		return nil, err
	}

	if !user.CheckPasswordHash(password) {
		r.logger.Error("invalid email or password")
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
