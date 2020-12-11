package gradeview_test

import (
	"reflect"
	"testing"

	"github.com/scottcalebc/gradeview"
)

var CatTests = []struct {
	name   string
	weight float32
}{
	{"Test", 1.0},
	{"Test", 2.0},
	{"Exam", 0.2},
	{"Quiz", 0.15},
}

func TestCategory(t *testing.T) {
	for _, tt := range CatTests {
		actual := gradeview.NewCategory(tt.name, tt.weight)

		if reflect.DeepEqual(*actual, tt) {
			t.Errorf("NewCategory(%s, %f): expected %+v, actual %+v", tt.name, tt.weight, tt, actual)
		}
	}
}
