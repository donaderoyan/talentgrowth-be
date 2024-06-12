package profile

type UpdateProfileInput struct {
	FirstName      string  `json:"firstName" validate:"required,alpha" updateValidation:"omitempty,alpha,min=3"`
	LastName       string  `json:"lastName" validate:"required,alpha" updateValidation:"omitempty,alpha"`
	Phone          string  `json:"phone" validate:"required,e164" updateValidation:"omitempty,e164"`
	Address        Address `json:"address" validate:"required" updateValidation:"omitempty"`
	Birthday       string  `json:"birthday" validate:"required,customdate,datebeforetoday" updateValidation:"omitempty,customdate,datebeforetoday"`
	Gender         string  `json:"gender" validate:"required,oneof=male female" updateValidation:"omitempty,min=3,oneof=male female"`
	Nationality    string  `json:"nationality" validate:"omitempty" updateValidation:"omitempty"`
	Bio            string  `json:"bio" validate:"omitempty" updateValidation:"omitempty"`
	ProfilePicture string  `json:"profilePicture" validate:"omitempty,url" updateValidation:"omitempty,url"`
}

type Address struct {
	Street     string `json:"street,omitempty" validate:"omitempty,min=3" updateValidation:"omitempty,min=3"`
	City       string `json:"city" validate:"required,min=3" updateValidation:"omitempty,min=3"`
	State      string `json:"state" validate:"required" updateValidation:"omitempty"`
	PostalCode string `json:"postalCode" validate:"required" updateValidation:"omitempty"`
	Country    string `json:"country" validate:"required" updateValidation:"omitempty"`
}
