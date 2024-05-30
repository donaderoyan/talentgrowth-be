package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// MusicalInfo struct for user's musical details
type MusicalInfo struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID               primitive.ObjectID `bson:"userID" json:"userID"` // Reference to the User
	SkillLevel           string             `bson:"skillLevel" json:"skillLevel"`
	PrimaryInstrument    string             `bson:"primaryInstrument" json:"primaryInstrument"`
	SecondaryInstruments []string           `bson:"secondaryInstruments,omitempty" json:"secondaryInstruments,omitempty"`
	Genres               []string           `bson:"genres" json:"genres"`
	FavoriteArtists      []string           `bson:"favoriteArtists,omitempty" json:"favoriteArtists,omitempty"`
	LearningGoals        []string           `bson:"learningGoals,omitempty" json:"learningGoals,omitempty"`
}
