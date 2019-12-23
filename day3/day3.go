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

// Part2 finds the intersection with the lowest combined signal delay
func Part2(args []string) int {
	wire0 := wires.NewWire(args[0])
	wire1 := wires.NewWire(args[1])

	intersections := wires.Intersections(wire0, wire1)

	signalDelays := map[wires.Point]int{}
	for _, point := range intersections {
		signalDelays[point] = wires.SignalDelay(point, wire0, wire1)
	}

	minDelay := math.MaxInt64
	for _, delay := range signalDelays {
		if delay < minDelay {
			minDelay = delay
		}
	}

	if minDelay < math.MaxInt64 {
		return minDelay
	}

	panic("no min delay found")
}
