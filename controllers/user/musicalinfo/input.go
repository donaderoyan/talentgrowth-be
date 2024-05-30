package musicalinfo

type MusicalInfoInput struct {
	SkillLevel           string   `json:"skill_level" updateValidation:"omitempty"`
	PrimaryInstrument    string   `json:"primary_instrument" updateValidation:"omitempty"`
	SecondaryInstruments []string `json:"secondary_instruments,omitempty" updateValidation:"omitempty"`
	Genres               []string `json:"genres" updateValidation:"omitempty"`
	FavoriteArtists      []string `json:"favorite_artists,omitempty" updateValidation:"omitempty"`
	LearningGoals        []string `json:"learning_goals,omitempty" updateValidation:"omitempty"`
}
