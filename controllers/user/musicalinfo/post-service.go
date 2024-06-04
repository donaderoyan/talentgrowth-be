package musicalinfo

import (
	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) CreateMusicalInfoService(userID string, input *MusicalInfoInput) (*model.MusicalInfo, error) {

	fields := bson.M{
		"skillLevel":           input.SkillLevel,
		"primaryInstrument":    input.PrimaryInstrument,
		"secondaryInstruments": input.SecondaryInstruments,
		"genres":               input.Genres,
		"favoriteArtists":      input.FavoriteArtists,
		"learningGoals":        input.LearningGoals,
	}

	result, err := s.repository.CreateMusicalInfo(userID, fields)

	return result, err
}
