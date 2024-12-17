package repository

import (
	"errors"
	"strings"

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
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			r.logger.Error(err.Error(), zap.Error(err))
			return err
		}
		r.logger.Error("Error inserting user", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetUserByUsernameAndPassword = `
SELECT id, username, password
FROM users
WHERE username=$1`

func (r *userRepository) GetUserByUsernameAndPassword(username, password string) (*entity.User, error) {

	user := &entity.User{Username: username}

	in := []interface{}{username}
	out := []interface{}{&user.Id, &user.Username, &user.Password}
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
SELECT id, username, password
FROM users
WHERE email=$1`

func (r *userRepository) GetUserByEmailAndPassword(email, password string) (*entity.User, error) {

	user := &entity.User{Email: email}

	in := []interface{}{email}
	out := []interface{}{&user.Id, &user.Username, &user.Password}
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

const QueryDeleteUser = `
DELETE FROM users
WHERE id = $1;`

func (r *userRepository) DeleteUser(userId entity.UUID) error {

	in := []interface{}{userId}
	if err := r.rdbms.Execute(QueryDeleteUser, in); err != nil {
		if err.Error() == rdbms.ErrNotFound {
			r.logger.Error("Error user not found", zap.Error(err))
			return err
		}

		r.logger.Error("Error find user by his id", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetOnlineUsers = `
SELECT 
    u.id,
    u.email,
    u.username,
    u.created_at
FROM 
    users u
JOIN 
    refresh_tokens rt ON u.id = rt.owner_id
WHERE 
    rt.last_refresh >= NOW() - INTERVAL '15 minutes'
LIMIT 5;`

func (r *userRepository) GetOnlineUsers() ([]entity.User, error) {

	limit := 5
	onlineUsers := make([]entity.User, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&onlineUsers[index].Id,
			&onlineUsers[index].Email,
			&onlineUsers[index].Username,
			&onlineUsers[index].CreatedAt,
		}
	}

	in := []any{}
	if err := r.rdbms.Query(QueryGetOnlineUsers, in, out); err != nil {
		r.logger.Error("Error retrieving onlineUsers", zap.Error(err))
		return []entity.User{}, err
	}

	return onlineUsers, nil
}
