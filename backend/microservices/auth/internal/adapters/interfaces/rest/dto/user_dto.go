package dto

import "github.com/Iamirup/whaler/backend/microservices/auth/internal/core/domain/entity"

// requests
type (
	RegisterRequest struct {
		Email           string `json:"email"            form:"email"            validate:"required,email"`
		Username        string `json:"username"         form:"username"         validate:"required,username"`
		Password        string `json:"password"         form:"password"         validate:"required,strong_password"`
		ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
	}

	LoginRequest struct {
		Email    string `json:"email,omitempty"    form:"email,omitempty"    validate:"omitempty,email"`
		Username string `json:"username,omitempty" form:"username,omitempty" validate:"omitempty,username"`
		Password string `json:"password"           form:"password"           validate:"required"`
	}

	AddAdminRequest struct {
		UserId entity.UUID `json:"user_id"    form:"user_id"    validate:"required"`
	}

	DeleteAdminRequest struct {
		AdminId entity.UUID `json:"admin_id"    form:"admin_id"    validate:"required"`
	}

	DeleteUserRequest struct {
		UserId entity.UUID `json:"user_id"    form:"user_id"    validate:"required"`
	}

	GetOnlineUsersRequest struct {
		// nothing (just cursor and limit in query parameters)
	}
)

// responses in successful status
type (
	RegisterResponse struct {
		// nothing
	}

	LoginResponse struct {
		// nothing
	}

	AddAdminResponse struct {
		// nothing
	}

	IsAdminResponse struct {
		IsAdmin bool `json:"is_admin"    form:"is_admin"`
	}

	DeleteAdminResponse struct {
		// nothing
	}

	DeleteUserResponse struct {
		// nothing
	}

	GetOnlineUsersResponse struct {
		OnlineUsers []entity.User `json:"online_users"     form:"online_users"`
	}
)

// responses in unsuccessful status
type (
	ErrorContent struct {
		Field   string `json:"field"     form:"field"`
		Message string `json:"message"   form:"message"`
	}

	ErrorResponse struct {
		Errors    []ErrorContent `json:"errors"       form:"errors"`
		NeedLogin bool           `json:"need_login"   form:"need_login"`
	}
)
