package shared

import "net/http"

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

var (
	ErrUnknownError = NewCustomError(http.StatusInternalServerError, "unknown error")

	ErrInvalidDateLayout      = NewCustomError(http.StatusBadRequest, "invalid date layout")
	ErrEmailAlreadyRegistered = NewCustomError(http.StatusBadRequest, "email already registered")
)
