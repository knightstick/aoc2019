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

func TestValueAt(t *testing.T) {
	t.Run("first value", func(t *testing.T) {
		program := intcode.NewProgram("1,0,0,0,99")
		actual := program.ValueAt(0)
		expected := 1

		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("second value", func(t *testing.T) {
		program := intcode.NewProgram("1,2,0,0,99")
		actual := program.ValueAt(1)
		expected := 2

		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("last value", func(t *testing.T) {
		program := intcode.NewProgram("1,2,0,0,99")
		actual := program.ValueAt(4)
		expected := 99

		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})
}
