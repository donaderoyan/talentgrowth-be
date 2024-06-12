package profile

import (
	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) PatchProfileService(userID string, input bson.M) (*model.User, error) {
	result, err := s.repository.PatchProfile(userID, input)
	return result, err
}
