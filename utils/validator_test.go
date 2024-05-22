package util_test

import (
	"testing"

	util "github.com/donaderoyan/talentgrowth-be/utils"

	model "github.com/donaderoyan/talentgrowth-be/models"
)

type testUser struct {
	FirstName      string        `json:"firstName" validate:"required,alpha" updateValidation:"omitempty,alpha"`
	LastName       string        `json:"lastName" validate:"required,alpha" updateValidation:"omitempty,alpha"`
	Phone          string        `json:"phone" validate:"required,e164" updateValidation:"omitempty,e164"`
	Address        model.Address `json:"address" validate:"omitempty" updateValidation:"omitempty"`
	Birthday       string        `json:"birthday" validate:"omitempty,customdate,datebeforetoday" updateValidation:"omitempty,customdate,datebeforetoday"`
	Gender         string        `json:"gender" validate:"omitempty,oneof=male female" updateValidation:"omitempty,oneof=male female"`
	Nationality    string        `json:"nationality" validate:"omitempty" updateValidation:"omitempty"`
	Bio            string        `json:"bio" validate:"omitempty" updateValidation:"omitempty"`
	ProfilePicture string        `json:"profilePicture" validate:"omitempty,url" updateValidation:"omitempty,url"`
}

func TestValidateLoginInput(t *testing.T) {
	input := testUser{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "+1234567890",
		Address:   model.Address{City: "City", State: "State", PostalCode: "12345", Country: "Country"},
		Birthday:  "01-01-1990",
		Gender:    "male",
	}

	err := util.Validator(input, "validate")
	if err != nil {
		t.Errorf("Validation failed for testUser: %v", err)
	}
}

func TestValidateUpdateProfileInput(t *testing.T) {
	input := testUser{
		FirstName: "Jane",
		LastName:  "Smith",
		Phone:     "+0987654321",
		Address:   model.Address{City: "New City", State: "New State", PostalCode: "54321", Country: "New Country"},
		Birthday:  "02-02-1992",
		Gender:    "female",
	}

	err := util.Validator(input, "updateValidation")
	if err != nil {
		t.Errorf("Validation failed for testUser: %v", err)
	}
}

func TestValidatePartialUpdateProfileInput(t *testing.T) {
	input := testUser{
		FirstName: "Jane",
		Phone:     "+0987654321",
		Address:   model.Address{State: "New State", PostalCode: "54321"},
	}

	err := util.Validator(input, "updateValidation")
	if err != nil {
		t.Errorf("Validation failed for testUser: %v", err)
	}
}
