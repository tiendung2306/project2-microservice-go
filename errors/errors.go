package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError, // 500
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    message,
	}
}

func NewUnauthorizedError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized, // 401
		Code:       "UNAUTHORIZED",
		Message:    message,
	}
}

func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound, // 404
		Code:       "NOT_FOUND",
		Message:    message,
	}
}

func NewConflictError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusConflict, // 409
		Code:       "CONFLICT",
		Message:    message,
	}
}
