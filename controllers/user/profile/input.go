package profileController

import model "github.com/donaderoyan/talentgrowth-be/models"

type UpdateProfileInput struct {
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
