package serr

import (
	"fmt"
)

type ServiceError struct {
	Message    string
	StatusCode int
	Details    any
}

func (e ServiceError) Error() string {
	return fmt.Sprintf(
		"%s - %v",
		e.Message, e.StatusCode,
	)
}

func ServiceErr(message string, statusCode int) error {
	return &ServiceError{
		Message:    message,
		StatusCode: statusCode,
	}
}
