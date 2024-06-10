package musicalinfo

import (
	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) UpdateMusicalInfoService(userID string, input bson.M) (*model.MusicalInfo, error) {
	result, err := s.repository.UpdateMusicalInfo(userID, input)
	return result, err
}
