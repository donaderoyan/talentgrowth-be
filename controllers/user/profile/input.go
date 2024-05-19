package profileController

type UpdateProfileInput struct {
	FirstName   string `json:"firstName" validate:"required,alpha"`
	LastName    string `json:"lastName" validate:"required,alpha"`
	Phone       string `json:"phone" validate:"required,e164"`
	Address     string `json:"address"`
	Birthday    string `json:"birthday" validate:"omitempty,customdate,datebeforetoday"`
	Gender      string `json:"gender" validate:"omitempty,oneof=male female"`
	Nationality string `json:"nationality"`
	Bio         string `json:"bio"`
}
