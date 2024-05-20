package profileController

import model "github.com/donaderoyan/talentgrowth-be/models"

type UpdateProfileInput struct {
	FirstName      string        `json:"firstName" validate:"required,alpha"`
	LastName       string        `json:"lastName" validate:"required,alpha"`
	Phone          string        `json:"phone" validate:"required,e164"`
	Address        model.Address `json:"address" validate:"omitempty"`
	Birthday       string        `json:"birthday" validate:"omitempty,customdate,datebeforetoday"`
	Gender         string        `json:"gender" validate:"omitempty,oneof=male female"`
	Nationality    string        `json:"nationality"`
	Bio            string        `json:"bio"`
	ProfilePicture string        `json:"profilePicture" validate:"omitempty,url"`
}
