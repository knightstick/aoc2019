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

func TestReplaceAt(t *testing.T) {
	t.Run("first value", func(t *testing.T) {
		program := intcode.NewProgram("1,0,0,0,99")
		program.ReplaceAt(1, 2)
		value := program.ValueAt(1)
		expected := 2

		if value != expected {
			t.Errorf("expected %d, got %d", expected, value)
		}
	})

	t.Run("last value", func(t *testing.T) {
		program := intcode.NewProgram("1,0,0,0,99,0")
		program.ReplaceAt(5, 99)
		value := program.ValueAt(5)
		expected := 99

		if value != expected {
			t.Errorf("expected %d, got %d", expected, value)
		}
	})
}

func TestExecutorConstruction(t *testing.T) {
	t.Run("add program", func(t *testing.T) {
		program := intcode.NewProgram("1,0,0,0,99")
		executor := intcode.NewExecutor(program)

		assertNextOpcode(t, executor, intcode.ADD)
	})

	t.Run("multiply program", func(t *testing.T) {
		program := intcode.NewProgram("2,0,0,0,99")
		executor := intcode.NewExecutor(program)

		assertNextOpcode(t, executor, intcode.MULTIPLY)
	})

	t.Run("halt program", func(t *testing.T) {
		program := intcode.NewProgram("99")
		executor := intcode.NewExecutor(program)

		assertNextOpcode(t, executor, intcode.HALT)
	})
}

func assertNextOpcode(t *testing.T, executor *intcode.Executor, opcode intcode.Opcode) {
	nextOpcode := executor.NextOpcode()

	if nextOpcode != opcode {
		t.Errorf("expected opcode ADD but got %d", opcode)
	}
}

func TestExecution(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		initial := intcode.NewProgram("1,0,0,0,99")
		actual := intcode.Execute(initial)
		expected := []int{2, 0, 0, 0, 99}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("multiplication", func(t *testing.T) {
		initial := intcode.NewProgram("2,3,0,3,99")
		actual := intcode.Execute(initial)
		expected := []int{2, 3, 0, 6, 99}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("multiplication again", func(t *testing.T) {
		initial := intcode.NewProgram("2,4,4,5,99,0")
		actual := intcode.Execute(initial)
		expected := []int{2, 4, 4, 5, 99, 9801}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("many steps", func(t *testing.T) {
		initial := intcode.NewProgram("1,1,1,4,99,5,6,0,99")
		actual := intcode.Execute(initial)
		expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
}
