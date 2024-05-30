package model

type CoursePreferences struct {
	PreferredLearningMode string   `bson:"preferredLearningMode" json:"preferredLearningMode"`
	Availability          []string `bson:"availability" json:"availability"`
	PreferredInstructors  []string `bson:"preferredInstructors,omitempty" json:"preferredInstructors,omitempty"` // this may change when Instructor model is implemented
}
