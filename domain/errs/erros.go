package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusNotFound,
	}
}

func NewUnExpectedError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusInternalServerError,
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
