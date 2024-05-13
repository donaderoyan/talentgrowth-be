package registerController

import model "github.com/donaderoyan/talentgrowth-be/models"

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
		Email:    input.Email,
		Password: input.Password,
	}

	result, err := s.repository.RegisterRepository(user)
	return result, err
}
