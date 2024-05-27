package login

import (
	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	LoginService(input *LoginInput) (*model.User, error)
	UpdateRememberTokenService(userID primitive.ObjectID, token string) error
}

type service struct {
	repository Repository
}

func NewLoginService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *LoginInput) (*model.User, error) {
	user := &model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	result, err := s.repository.LoginRepository(user)
	return result, err
}

func (s *service) UpdateRememberTokenService(userID primitive.ObjectID, token string) error {
	return s.repository.UpdateRememberTokenRepository(userID, token)
}
