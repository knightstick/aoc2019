package main

import (
	"testing"

	"github.com/knightstick/aoc2019/manhattandistance"
	"github.com/knightstick/aoc2019/wires"
)

func TestManhattanDistance(t *testing.T) {
	t.Run("origin", func(t *testing.T) {
		assertManhattanDistance(t, wires.Point{X: 0, Y: 0}, 0)
	})

	t.Run("(1,0)", func(t *testing.T) {
		assertManhattanDistance(t, wires.Point{X: 1, Y: 0}, 1)
	})

	t.Run("(0,1)", func(t *testing.T) {
		assertManhattanDistance(t, wires.Point{X: 0, Y: 1}, 1)
	})

	t.Run("(1,1)", func(t *testing.T) {
		assertManhattanDistance(t, wires.Point{X: 1, Y: 1}, 2)
	})

	t.Run("(2,3)", func(t *testing.T) {
		assertManhattanDistance(t, wires.Point{X: 2, Y: 3}, 5)
	})

	t.Run("(-3,5)", func(t *testing.T) {
		assertManhattanDistance(t, wires.Point{X: -3, Y: 5}, 8)
	})

	t.Run("(-7,-11)", func(t *testing.T) {
		assertManhattanDistance(t, wires.Point{X: -7, Y: -11}, 18)
	})
}

func assertManhattanDistance(t *testing.T, point wires.Point, expected int) {
	actual := manhattandistance.FromOrigin(point)

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
