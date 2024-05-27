package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// MusicalInformation struct for user's musical details
type MusicalInformation struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID               primitive.ObjectID `bson:"user_id" json:"user_id"` // Reference to the User
	SkillLevel           string             `bson:"skill_level" json:"skill_level"`
	PrimaryInstrument    string             `bson:"primary_instrument" json:"primary_instrument"`
	SecondaryInstruments []string           `bson:"secondary_instruments,omitempty" json:"secondary_instruments,omitempty"`
	Genres               []string           `bson:"genres" json:"genres"`
	FavoriteArtists      []string           `bson:"favorite_artists,omitempty" json:"favorite_artists,omitempty"`
	LearningGoals        []string           `bson:"learning_goals,omitempty" json:"learning_goals,omitempty"`
}
