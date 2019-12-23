package wires

import (
	"reflect"
	"strconv"
	"strings"
)

// Wire a list of points representing the segments of wire on the grid
type Wire struct {
	Segments []Point
}

// Point an X Y pair representing a point on the grid
type Point struct {
	X int
	Y int
}

// Origin is the point (0, 0) where the wires start
var Origin = Point{X: 0, Y: 0}

// NewWire builds a new Wire
func NewWire(input string) *Wire {
	wire := new(Wire)
	wire.Segments = segmentsFromString(input)
	return wire
}

// Intersections finds all points where the wires cross each other
func Intersections(wire0, wire1 *Wire) (result []Point) {
	hashmap := map[Point]bool{}

	for _, point := range wire0.Segments {
		if !hashmap[point] {
			hashmap[point] = true
		}
	}

	for _, point := range wire1.Segments {
		if !isOrigin(point) {
			if hashmap[point] {
				result = append(result, point)
			}
		}
	}

	return
}

func SignalDelay(point Point, wire0, wire1 *Wire) int {
	return wireDistance(point, wire0) + wireDistance(point, wire1)
}

func wireDistance(point Point, wire *Wire) (result int) {
	for index, other := range wire.Segments {
		if reflect.DeepEqual(point, other) {
			return index
		}
	}
	panic("point not in wire")
}

type direction struct {
	cardinality cardinality
	distance    int
}

type cardinality rune

const (
	// CARDR is the cardinality in the Right direction
	CARDR = 'R'
	// CARDL is the cardinality in the Left direction
	CARDL = 'L'
	// CARDU is the cardinality in the Up direction
	CARDU = 'U'
	// CARDD is the cardinality in the Down direction
	CARDD = 'D'
)

var cardinalities = map[cardinality]bool{
	CARDR: true,
	CARDL: true,
	CARDU: true,
	CARDD: true,
}

func segmentsFromString(input string) []Point {
	points := []Point{}
	// Add the origin
	points = append(points, Point{0, 0})

	if input == "" {
		return points
	}

	directionStrings := strings.Split(input, ",")
	directions := []*direction{}

	for _, directionString := range directionStrings {
		directions = append(directions, newDirection(directionString))
	}

	for _, direction := range directions {
		points = append(points, pointsForDirection(points, direction)...)
	}

	return points
}

func newDirection(input string) *direction {
	cardinality := cardinality(input[0])

	if !cardinalities[cardinality] {
		panic("unknown cardinality")
	}

	distance, err := strconv.Atoi(input[1:])

	if err != nil {
		panic("unknown direction distance")
	}

	return &direction{
		cardinality: cardinality,
		distance:    distance,
	}
}

func pointsForDirection(points []Point, direction *direction) (result []Point) {
	origin := points[len(points)-1]

	for index := 0; index < direction.distance; index++ {
		nextPoint := nextPoint(origin, direction.cardinality)
		result = append(result, nextPoint)
		origin = nextPoint
	}

	return
}

func nextPoint(origin Point, cardinality cardinality) Point {
	switch cardinality {
	case CARDR:
		return Point{X: origin.X + 1, Y: origin.Y}
	case CARDL:
		return Point{X: origin.X - 1, Y: origin.Y}
	case CARDU:
		return Point{X: origin.X, Y: origin.Y + 1}
	case CARDD:
		return Point{X: origin.X, Y: origin.Y - 1}
	}
	return Point{}
}

func pointInWire(point Point, wire *Wire) bool {
	for _, other := range wire.Segments {
		if reflect.DeepEqual(point, other) {
			return true
		}
	}
	return false
}

func isOrigin(point Point) bool {
	return reflect.DeepEqual(point, Origin)
}
