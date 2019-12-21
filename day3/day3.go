package day3

import (
	"math"

	"github.com/knightstick/aoc2019/manhattandistance"
	"github.com/knightstick/aoc2019/wires"
)

// Part1 finds the nearest intersection of the two wires coming out of the
// fuel management system
func Part1(args []string) int {
	wire0 := wires.NewWire(args[0])
	wire1 := wires.NewWire(args[1])

	intersections := wires.Intersections(wire0, wire1)

	distances := map[wires.Point]int{}
	for _, point := range intersections {
		distances[point] = manhattandistance.FromOrigin(point)
	}

	minDistance := math.MaxInt64
	for _, distance := range distances {
		if distance < minDistance {
			minDistance = distance
		}
	}

	if minDistance < math.MaxInt64 {
		return minDistance
	}

	panic("no min distance found")
}
