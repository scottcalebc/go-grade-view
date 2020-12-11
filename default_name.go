package gradeview

// Category holds information for various categories in a class such as
// 		Assignments, Exams, Quizzes, Extra Credit, etc.
type Category struct {
	name   string
	weight float32
}

// Create Category struct without accessing unexported fields
func NewCategory(name string, weight float32) *Category {
	p := Category{name: name, weight: weight}

	return &p
}
