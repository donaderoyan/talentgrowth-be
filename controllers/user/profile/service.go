package profile

import (
	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	PatchProfileService(userID string, input bson.M) (*model.User, error)
	PutProfileService(userID string, input *UpdateProfileInput) (*model.User, error)
	GetProfileService(userID string) (*model.User, error)
}

type service struct {
	repository Repository
}

func NewProfileService(repository Repository) *service {
	return &service{repository: repository}
}
