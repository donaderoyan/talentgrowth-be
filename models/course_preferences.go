package model

type CoursePreferences struct {
	PreferredLearningMode string   `bson:"preferred_learning_mode" json:"preferred_learning_mode"`
	Availability          []string `bson:"availability" json:"availability"`
	PreferredInstructors  []string `bson:"preferred_instructors,omitempty" json:"preferred_instructors,omitempty"` // this may change when Instructor model is implemented
}
