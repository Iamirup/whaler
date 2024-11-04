package dto

type (
	RegisterRequest struct {
		Email           string `json:"email" validate:"required,email"`
		Username        string `json:"username" validate:"required,regexp=^[A-Za-z][A-Za-z0-9_]{2,31}$"`
		Password        string `json:"password" validate:"required,regexp=^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\\$%\\^&\\*]).{8,255}$"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"omitempty,email"`
		Username string `json:"username" validate:"omitempty,regexp=^[A-Za-z][A-Za-z0-9_]{2,31}$"`
		Password string `json:"password" validate:"required"`
	}
)
