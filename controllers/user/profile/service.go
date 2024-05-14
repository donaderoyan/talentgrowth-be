package profileController

import (
	"fmt"
	"time"

	model "github.com/donaderoyan/talentgrowth-be/models"
)

type Service interface {
	UpdateProfileService(userID string, input *UpdateProfileInput) (*model.User, error)
}

type service struct {
	repository Repository
}

func NewProfileService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateProfileService(userID string, input *UpdateProfileInput) (*model.User, error) {
	// Update user profile
	parsedBirthday, _ := time.Parse("02-01-2006", input.Birthday)
	user := &model.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		Address:     input.Address,
		Birthday:    parsedBirthday,
		Gender:      input.Gender,
		Nationality: input.Nationality,
		Bio:         input.Bio,
	}

	fmt.Printf("Input Birthday: %s\n", input.Birthday)
	fmt.Printf("Parsed Birthday: %v\n", parsedBirthday)

	result, err := s.repository.UpdateProfile(userID, user)
	return result, err
}
