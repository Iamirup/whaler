package dto

type (
	RegisterRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,username"`
		Password string `json:"password" validate:"required,strong_password"`
		Confirmp string `json:"confirmp" validate:"required,eqfield=Password"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"omitempty,email"`
		Username string `json:"username" validate:"omitempty,username"`
		Password string `json:"password" validate:"required"`
	}
)
