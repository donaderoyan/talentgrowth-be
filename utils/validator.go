package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// Validator function to validate struct fields based on tags
func Validator(s interface{}, tagName string) error {
	// Check if the input is a struct
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return fmt.Errorf("validation error: expected a struct, got %s", reflect.TypeOf(s).Kind())
	}

	// Initialize the validator
	validate := validator.New()
	if tagName != "" {
		validate.SetTagName(tagName)
	}

	// Register custom date validation for dd-mm-yyyy format
	validate.RegisterValidation("customdate", func(fl validator.FieldLevel) bool {
		const layout = "02-01-2006" // dd-mm-yyyy
		_, err := time.Parse(layout, fl.Field().String())
		return err == nil
	})

	// Register custom validation to check if the date is before today's date
	validate.RegisterValidation("datebeforetoday", func(fl validator.FieldLevel) bool {
		const layout = "02-01-2006" // dd-mm-yyyy
		date, err := time.Parse(layout, fl.Field().String())
		if err != nil {
			return false
		}
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		return date.Before(today)
	})

	if err := validate.Struct(s); err != nil {
		// Handle invalid argument error
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Errorf("invalid argument for validation: %v", err)
		}

		// Collect all field validation errors
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			tag := err.ActualTag()
			field := err.Field()
			param := err.Param()

			// Use msgForTag to get the error message for each validation error
			message := msgForTag(tag, param)
			if message == "" {
				message = fmt.Sprintf("Please ensure the %s field meets the requirement: %s", field, tag)
			}
			errorMessages = append(errorMessages, fmt.Sprintf("Please ensure the %s field meets the requirement: %s", field, message))
		}
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}

func msgForTag(tag string, param any) string {
	switch tag {
	case "required":
		return "This field is required."
	case "email":
		return "This field must contain a valid email address."
	case "alpha":
		return "This field must contain only alphabetic characters."
	case "e164":
		return "This field must contain a valid phone number format, e.g., +12345678900."
	case "min":
		if param != "" {
			return fmt.Sprintf("This field must contain at least %s characters.", param)
		}
	case "max":
		if param != "" {
			return fmt.Sprintf("This field must contain at most %s characters.", param)
		}
	case "oneof":
		if param == "male female" {
			return "This field must be either 'male' or 'female'."
		}
		return "This field must be one of the specified values."
	case "customdate":
		return "This field must be in dd-mm-yyyy format."
	case "datebeforetoday":
		return "This field must be before today's date."

	}
	return ""
}

func InterfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = v.(string)
	}
	return result
}
