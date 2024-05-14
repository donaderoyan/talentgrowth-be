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
			if err.Param() != "" || err.ActualTag() == "e164" {
				val1 := err.ActualTag()
				val2 := err.Param()
				if strings.ToLower(err.Field()) == "phone" {
					val1 = "format"
					val2 = "a phone number in international, like +12345678900"
				}
				if strings.ToLower(err.ActualTag()) == "min" || strings.ToLower(err.ActualTag()) == "max" {
					val2 = err.Param() + " characters"
				}
				errorMessages = append(errorMessages, fmt.Sprintf("Please ensure the %s field meets the requirement: %s should be %s.", err.Field(), val1, val2))
			} else {
				errorMessages = append(errorMessages, fmt.Sprintf("Please ensure the %s field meets the requirement: %s.", err.Field(), err.ActualTag()))
			}
		}
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}
