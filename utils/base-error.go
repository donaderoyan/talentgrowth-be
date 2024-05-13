package util

import (
	"fmt"
)

// BaseError defines a standard error structure for the application.
type BaseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface for BaseError.
func (e *BaseError) Error() string {
	return fmt.Sprintf("Error: %s - %s", e.Code, e.Message)
}

// NewBaseError creates a new instance of BaseError.
func NewBaseError(code, message string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
	}
}
