package errors

import (
	"fmt"
	"net/http"
)

// AppError là một cấu trúc lỗi tùy chỉnh cho ứng dụng
type AppError struct {
	StatusCode int    `json:"-"`       // HTTP status code
	Code       string `json:"code"`    // Mã lỗi nội bộ
	Message    string `json:"message"` // Thông điệp lỗi cho người dùng
	Detail     string `json:"detail"`  // Thông tin chi tiết thêm về lỗi (tùy chọn)
	Err        error  `json:"-"`       // Lỗi gốc (nếu có)
}

// Error implement interface error
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap unwraps the error to its original error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewBadRequestError tạo một lỗi Bad Request (400)
func NewBadRequestError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Code:       "BAD_REQUEST",
		Message:    message,
		Err:        err,
	}
}

// NewUnauthorizedError tạo một lỗi Unauthorized (401)
func NewUnauthorizedError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       "UNAUTHORIZED",
		Message:    message,
		Err:        err,
	}
}

// NewForbiddenError tạo một lỗi Forbidden (403)
func NewForbiddenError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusForbidden,
		Code:       "FORBIDDEN",
		Message:    message,
		Err:        err,
	}
}

// NewNotFoundError tạo một lỗi Not Found (404)
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Code:       "NOT_FOUND",
		Message:    message,
		Err:        err,
	}
}

// NewConflictError tạo một lỗi Conflict (409)
func NewConflictError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusConflict,
		Code:       "CONFLICT",
		Message:    message,
		Err:        err,
	}
}

// NewInternalServerError tạo một lỗi Internal Server Error (500)
func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    message,
		Err:        err,
	}
}

// IsAppError kiểm tra xem một lỗi có phải là AppError
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// WithDetail thêm thông tin chi tiết vào lỗi
func (e *AppError) WithDetail(detail string) *AppError {
	e.Detail = detail
	return e
}
