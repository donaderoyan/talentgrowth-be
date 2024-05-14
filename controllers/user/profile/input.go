package profileController

import (
	"fmt"
	"time"
)

type UpdateProfileInput struct {
	FirstName   string `json:"firstName" validate:"required,alpha"`
	LastName    string `json:"lastName" validate:"required,alpha"`
	Phone       string `json:"phone" validate:"required,e164"`
	Address     string `json:"address"`
	Birthday    string `json:"birthday" validate:"omitempty"`
	Gender      string `json:"gender" validate:"omitempty,oneof=male female other"`
	Nationality string `json:"nationality"`
	Bio         string `json:"bio"`
}

func ValidateBirthday(birthday string) (bool, error) {
	if birthday == "" {
		return true, nil // Empty is allowed due to 'omitempty'
	}

	const formatLayout = "05-04-2006" // dd-mm-yyyy
	_, err := time.Parse(formatLayout, birthday)
	if err != nil {
		return false, fmt.Errorf("invalid date format for Birthday, expected dd-mm-yyyy: %v", err)
	}
	return true, nil
}

// func validateCustomDate(fl validator.FieldLevel) bool {
// 	const formatLayout = "02-01-2006" // dd-mm-yyyy
// 	dateStr := fl.Field().String()
// 	date, err := time.Parse(formatLayout, dateStr)
// 	if err != nil {
// 		fmt.Printf("Error parsing date: %s with format: %s\n", dateStr, formatLayout)
// 		return false
// 	}
// 	isValid := date.Before(time.Now())
// 	fmt.Printf("Date validation for %s: %t\n", dateStr, isValid)
// 	return isValid
// }

// var validate *validator.Validate

// func init() {
// 	validate := validator.New()
// 	customDateValidator := func(fl validator.FieldLevel) bool {
// 		return validateCustomDate(fl)
// 	}
// 	if err := validate.RegisterValidation("customDate", customDateValidator); err != nil {
// 		fmt.Printf("Error registering customDate validator: %v\n", err)
// 	} else {
// 		fmt.Println("Custom validation for 'customDate' registered successfully")
// 	}
// }
