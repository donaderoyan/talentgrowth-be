package profile

import model "github.com/donaderoyan/talentgrowth-be/models"

func (s *service) GetProfileService(userID string) (*model.User, error) {
	result, err := s.repository.GetProfile(userID)
	return result, err
}
