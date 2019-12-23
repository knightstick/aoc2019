package main

import (
	"reflect"
	"testing"

	"github.com/knightstick/aoc2019/wires"
)

func TestWireConstructor(t *testing.T) {
	t.Run("zero points", func(t *testing.T) {
		actual := wires.NewWire("")
		expected := &wires.Wire{
			Segments: []wires.Point{wires.Point{X: 0, Y: 0}},
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("one point right", func(t *testing.T) {
		actual := wires.NewWire("R1")
		expected := &wires.Wire{
			Segments: []wires.Point{
				wires.Point{X: 0, Y: 0},
				wires.Point{X: 1, Y: 0},
			},
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("three points right", func(t *testing.T) {
		actual := wires.NewWire("R3")
		expected := &wires.Wire{
			Segments: []wires.Point{
				wires.Point{X: 0, Y: 0},
				wires.Point{X: 1, Y: 0},
				wires.Point{X: 2, Y: 0},
				wires.Point{X: 3, Y: 0},
			},
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("three points left", func(t *testing.T) {
		actual := wires.NewWire("L3")
		expected := &wires.Wire{
			Segments: []wires.Point{
				wires.Point{X: 0, Y: 0},
				wires.Point{X: -1, Y: 0},
				wires.Point{X: -2, Y: 0},
				wires.Point{X: -3, Y: 0},
			},
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("three points up", func(t *testing.T) {
		actual := wires.NewWire("U3")
		expected := &wires.Wire{
			Segments: []wires.Point{
				wires.Point{X: 0, Y: 0},
				wires.Point{X: 0, Y: 1},
				wires.Point{X: 0, Y: 2},
				wires.Point{X: 0, Y: 3},
			},
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("three points down", func(t *testing.T) {
		actual := wires.NewWire("D3")
		expected := &wires.Wire{
			Segments: []wires.Point{
				wires.Point{X: 0, Y: 0},
				wires.Point{X: 0, Y: -1},
				wires.Point{X: 0, Y: -2},
				wires.Point{X: 0, Y: -3},
			},
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("all directions", func(t *testing.T) {
		actual := wires.NewWire("R2,U2,L2,D1")
		expected := &wires.Wire{
			Segments: []wires.Point{
				wires.Point{X: 0, Y: 0},
				wires.Point{X: 1, Y: 0},
				wires.Point{X: 2, Y: 0},
				wires.Point{X: 2, Y: 1},
				wires.Point{X: 2, Y: 2},
				wires.Point{X: 1, Y: 2},
				wires.Point{X: 0, Y: 2},
				wires.Point{X: 0, Y: 1},
			},
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
}

func TestIntersections(t *testing.T) {
	t.Run("one crossing", func(t *testing.T) {
		assertIntersections(
			t, wires.NewWire("R2,U3"), wires.NewWire("U2,R3"),
			[]wires.Point{wires.Point{X: 2, Y: 2}})
	})

	t.Run("three crossings", func(t *testing.T) {
		assertIntersections(
			t, wires.NewWire("R1,D3"), wires.NewWire("D1,R3,D1,L3,D1,R3"),
			[]wires.Point{
				wires.Point{X: 1, Y: -1},
				wires.Point{X: 1, Y: -2},
				wires.Point{X: 1, Y: -3},
			})
	})
}

func assertIntersections(t *testing.T, wire0, wire1 *wires.Wire, expected []wires.Point) {
	actual := wires.Intersections(wire0, wire1)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestSignalDelay(t *testing.T) {
	t.Run("first point on both", func(t *testing.T) {
		wire0 := wires.NewWire("R1")
		wire1 := wires.NewWire("R1")
		actual := wires.SignalDelay(wires.Point{X: 1, Y: 0}, wire0, wire1)
		expected := 2

		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("a few steps on both", func(t *testing.T) {
		wire0 := wires.NewWire("R1,U2,R1,U2")
		wire1 := wires.NewWire("U3,L1,U1,R2,D1,R2")
		actual := wires.SignalDelay(wires.Point{X: 2, Y: 3}, wire0, wire1)
		expected := 5 + 9

		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})
}
