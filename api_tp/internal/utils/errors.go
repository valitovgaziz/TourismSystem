// utils/errors.go
package utils

import (
	"fmt"
	"net/http"
)

// APIError представляет ошибку API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d: %s", e.Code, e.Message)
}

// ErrorResponse представляет стандартный ответ с ошибкой
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ValidationErrorResponse представляет ответ с ошибками валидации
type ValidationErrorResponse struct {
	Error   bool              `json:"error"`
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// Predefined errors
var (
	ErrInvalidJSON     = &APIError{Code: http.StatusBadRequest, Message: "Invalid JSON"}
	ErrEmptyRequestBody = &APIError{Code: http.StatusBadRequest, Message: "Request body is empty"}
	ErrRequestBodyTooLarge = &APIError{Code: http.StatusRequestEntityTooLarge, Message: "Request body too large"}
)