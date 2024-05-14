package registerController

import (
	"time"

	model "github.com/donaderoyan/talentgrowth-be/models"
)

type Service interface {
	RegisterService(input *RegisterInput) (*model.User, error)
}

type service struct {
	repository Repository
}

func NewRegisterService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *RegisterInput) (*model.User, error) {
	user := &model.User{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := s.repository.RegisterRepository(user)
	return result, err
}
