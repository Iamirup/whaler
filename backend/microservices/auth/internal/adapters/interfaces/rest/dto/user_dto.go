package dto

type (
	RegisterRequest struct {
		Email           string `json:"email" validate:"required"`
		Username        string `json:"username" validate:"required"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password"`
	}
)
