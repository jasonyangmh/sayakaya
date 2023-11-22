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

	ErrInvalidID              = NewCustomError(http.StatusBadRequest, "invalid id")
	ErrInvalidDateLayout      = NewCustomError(http.StatusBadRequest, "invalid date layout")
	ErrEmailAlreadyRegistered = NewCustomError(http.StatusBadRequest, "email already registered")
	ErrInvalidPromo           = NewCustomError(http.StatusBadRequest, "invalid promo")
	ErrPromoIsUsed            = NewCustomError(http.StatusBadRequest, "promo is used")
	ErrPromoHasEnded          = NewCustomError(http.StatusBadRequest, "promo has ended")
	ErrUserIsNotVerified      = NewCustomError(http.StatusBadRequest, "user is not verified")
)
