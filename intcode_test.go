package main

import (
	"reflect"
	"testing"

	"github.com/knightstick/aoc2019/intcode"
)

func TestProgramConstruction(t *testing.T) {
	var conconstructorTests = []struct {
		String  string
		Program intcode.Program
	}{
		{"1,0,0,0,99", intcode.Program{Values: []int{1, 0, 0, 0, 99}}},
		{"2,3,0,3,99", intcode.Program{Values: []int{2, 3, 0, 3, 99}}},
		{"2,4,4,5,99,0", intcode.Program{Values: []int{2, 4, 4, 5, 99, 0}}},
	}

	for _, tt := range conconstructorTests {
		actual := intcode.NewProgram(tt.String)
		expected := &tt.Program

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}
