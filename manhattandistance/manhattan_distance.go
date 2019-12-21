package manhattandistance

import "github.com/knightstick/aoc2019/wires"

import "math"

// FromOrigin returns the distance from the Point (0, 0)
func FromOrigin(point wires.Point) int {
	xComponent := int(math.Abs(float64(point.X)))
	yComponent := int(math.Abs(float64(point.Y)))
	return xComponent + yComponent
}
