package register

type RegisterInput struct {
	FirstName string `json:"firstName" validate:"required,alpha"`
	LastName  string `json:"lastName" validate:"required,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}
