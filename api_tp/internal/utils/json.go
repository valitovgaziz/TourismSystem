// utils/json.go
package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// DecodeJSON декодирует JSON из тела запроса с валидацией
func DecodeJSON(r *http.Request, v interface{}) error {
	// Ограничиваем размер тела запроса (например, 1MB)
	maxBytes := int64(1_048_576) // 1MB
	r.Body = http.MaxBytesReader(nil, r.Body, maxBytes)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Запрещаем неизвестные поля

	err := decoder.Decode(v)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case err == io.EOF:
			return &APIError{
				Code:    http.StatusBadRequest,
				Message: "Request body is empty",
			}
		case err.Error() == "http: request body too large":
			return &APIError{
				Code:    http.StatusRequestEntityTooLarge,
				Message: fmt.Sprintf("Request body must not be larger than %d bytes", maxBytes),
			}
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return &APIError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("Unknown field in JSON: %s", fieldName),
			}
		case errors.As(err, &syntaxError):
			return &APIError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("Malformed JSON at position %d", syntaxError.Offset),
			}
		case errors.As(err, &unmarshalTypeError):
			return &APIError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("Invalid value for field '%s'. Expected type %s", unmarshalTypeError.Field, unmarshalTypeError.Type),
			}
		case errors.As(err, &invalidUnmarshalError):
			return &APIError{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
			}
		default:
			return &APIError{
				Code:    http.StatusBadRequest,
				Message: "Invalid JSON",
			}
		}
	}

	// Проверяем, что нет лишних данных после JSON
	if err = decoder.Decode(&struct{}{}); err != io.EOF {
		return &APIError{
			Code:    http.StatusBadRequest,
			Message: "Request body must contain only single JSON object",
		}
	}

	return nil
}

// WriteJSON записывает JSON ответ
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return nil
	}

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true) // Экранируем HTML для безопасности

	return encoder.Encode(data)
}

// WriteError записывает ошибку в формате JSON
func WriteError(w http.ResponseWriter, status int, message string) {
	errorResponse := ErrorResponse{
		Error:   true,
		Message: message,
		Code:    status,
	}

	WriteJSON(w, status, errorResponse)
}

// WriteValidationError записывает ошибки валидации
func WriteValidationError(w http.ResponseWriter, errors map[string]string) {
	errorResponse := ValidationErrorResponse{
		Error:   true,
		Message: "Validation failed",
		Code:    http.StatusBadRequest,
		Errors:  errors,
	}

	WriteJSON(w, http.StatusBadRequest, errorResponse)
}