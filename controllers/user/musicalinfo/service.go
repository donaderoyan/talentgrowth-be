package musicalinfo

import (
	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	UpdateMusicalInfoService(userID string, input bson.M) (*model.MusicalInfo, error)
	CreateMusicalInfoService(userID string, input *MusicalInfoInput) (*model.MusicalInfo, error)
}

type service struct {
	repository Repository
}

func NewMusicalInfoService(repository Repository) *service {
	return &service{repository: repository}
}
