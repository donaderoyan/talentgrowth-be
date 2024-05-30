package model

type Model struct {
	Model interface{}
}

func RegisterModels() []string {
	return []string{
		"user",
		"musicalinfo",
		"coursepreferences",
		// "Product", etc ...
	}
}
