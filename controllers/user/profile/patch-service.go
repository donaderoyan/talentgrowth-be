package profile

import (
	"reflect"
	"time"

	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) UpdateProfileService(userID string, input *UpdateProfileInput) (*model.User, error) {
	// Update user profile
	// Ensure correct field names are used for MongoDB document
	correctFieldNames := bson.M{
		"firstName":      input.FirstName,
		"lastName":       input.LastName,
		"birthday":       input.Birthday,
		"phone":          input.Phone,
		"gender":         input.Gender,
		"nationality":    input.Nationality,
		"bio":            input.Bio,
		"updatedAt":      time.Now(),
		"profilePicture": input.ProfilePicture,
	}

	// Handle address update if present
	if input.Address != (Address{}) {
		correctFieldNames["address.street"] = input.Address.Street
		correctFieldNames["address.city"] = input.Address.City
		correctFieldNames["address.state"] = input.Address.State
		correctFieldNames["address.postalCode"] = input.Address.PostalCode
		correctFieldNames["address.country"] = input.Address.Country
	}

	// Remove empty or nil fields to avoid overwriting with zero values
	for key, value := range correctFieldNames {
		if value == nil || (reflect.TypeOf(value).Kind() == reflect.String && value == "") {
			delete(correctFieldNames, key)
		} else if val, ok := value.(Address); ok && val == (Address{}) {
			delete(correctFieldNames, key)
		}
	}

	if input.Birthday != "" {
		var parsedBirthday time.Time
		var err error
		parsedBirthday, err = time.Parse("02-01-2006", input.Birthday)
		if err != nil {
			return nil, err
		}
		correctFieldNames["birthday"] = parsedBirthday
	}

	result, err := s.repository.UpdateProfile(userID, correctFieldNames)
	return result, err
}