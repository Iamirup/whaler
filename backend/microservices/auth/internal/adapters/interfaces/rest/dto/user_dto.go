package dto

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
)

// responses in successful status
type (
	RegisterResponse struct {
		AccessToken string `json:"access_token"  form:"access_token"`
	}

	LoginResponse struct {
		AccessToken string `json:"access_token"  form:"access_token"`
	}
)

// responses in unsuccessful status
type (
	ErrorContent struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ErrorResponse struct {
		Errors    []ErrorContent `json:"errors"`
		NeedLogin bool           `json:"need_login"`
	}
)
