package profile

import model "github.com/donaderoyan/talentgrowth-be/models"

type Service interface {
	UpdateProfileService(userID string, input *UpdateProfileInput) (*model.User, error)
	PutProfileService(userID string, input *UpdateProfileInput) (*model.User, error)
	GetProfileService(userID string) (*model.User, error)
}

type service struct {
	repository Repository
}

func NewProfileService(repository Repository) *service {
	return &service{repository: repository}
}
