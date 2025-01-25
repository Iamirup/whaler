package dto

import "github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/domain/entity"

// requests
type (
	UpdateTableConfigRequest struct {
		TableConfig entity.TableConfig `json:"table_config"  form:"table_config"`
	}

	SeeTableRequest struct {
		// nothing (just cursor and limit in query parameters)
	}
)

// responses in successful status
type (
	UpdateTableConfigResponse struct {
		// nothing
	}

	SeeTableResponse struct {
		Table     []entity.TableRecord `json:"table"         form:"table"`
		NewCursor string               `json:"new_cursor"   form:"new_cursor"`
	}
)

// responses in unsuccessful status
type (
	ErrorContent struct {
		Field   string `json:"field"     form:"field"`
		Message string `json:"message"   form:"message"`
	}

	ErrorResponse struct {
		Errors      []ErrorContent `json:"errors"         form:"errors"`
		NeedRefresh bool           `json:"need_refresh"   form:"need_refresh"`
	}
)
