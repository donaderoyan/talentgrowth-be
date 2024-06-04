package musicalinfo

type MusicalInfoInput struct {
	SkillLevel           string   `json:"skill_level" validate:"required" updateValidation:"omitempty"`
	PrimaryInstrument    string   `json:"primary_instrument" validate:"required" updateValidation:"omitempty"`
	SecondaryInstruments []string `json:"secondary_instruments,omitempty" validate:"omitempty" updateValidation:"omitempty"`
	Genres               []string `json:"genres" validate:"omitempty" updateValidation:"omitempty"`
	FavoriteArtists      []string `json:"favorite_artists,omitempty" validate:"omitempty" updateValidation:"omitempty"`
	LearningGoals        []string `json:"learning_goals,omitempty" validate:"omitempty" updateValidation:"omitempty"`
}
