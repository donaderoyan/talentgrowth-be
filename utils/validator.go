package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator function to validate struct fields based on tags
func Validator(s interface{}) error {
	// Check if the input is a struct
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return fmt.Errorf("validation error: expected a struct, got %s", reflect.TypeOf(s).Kind())
	}

	// Initialize the validator
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		// Handle invalid argument error
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Errorf("invalid argument for validation: %v", err)
		}

		// Collect all field validation errors
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("The field '%s' is invalid: it must satisfy the condition '%s'.", err.Field(), err.ActualTag()))
		}
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}
