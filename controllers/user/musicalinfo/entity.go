package musicalinfo

type MusicalInfoInput struct {
	SkillLevel           string   `json:"skillLevel" validate:"required,alpha" updateValidation:"omitempty"`
	PrimaryInstrument    string   `json:"primaryInstrument" validate:"required,alpha" updateValidation:"omitempty"`
	SecondaryInstruments []string `json:"secondaryInstruments,omitempty" validate:"omitempty" updateValidation:"omitempty"`
	Genres               []string `json:"genres" validate:"omitempty" updateValidation:"omitempty"`
	FavoriteArtists      []string `json:"favoriteArtists,omitempty" validate:"omitempty" updateValidation:"omitempty"`
	LearningGoals        []string `json:"learningGoals,omitempty" validate:"omitempty" updateValidation:"omitempty"`
}
