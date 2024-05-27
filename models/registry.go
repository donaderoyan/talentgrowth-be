package model

type Model struct {
	Model interface{}
}

func RegisterModels() []string {
	return []string{
		"User",
		"MusicalInformation",
		"CoursePreferences",
		// "Product", etc ...
	}
}
