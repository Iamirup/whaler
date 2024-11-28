package dto

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
