package musicalinfo

import (
	model "github.com/donaderoyan/talentgrowth-be/models"
)

type Service interface {
	UpdateMusicalInfoService(userID string, input *MusicalInfoInput) (*model.MusicalInfo, error)
	CreateMusicalInfoService(userID string, input *MusicalInfoInput) (*model.MusicalInfo, error)
}

type service struct {
	repository Repository
}

func NewMusicalInfoService(repository Repository) *service {
	return &service{repository: repository}
}
