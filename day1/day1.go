package day1

import (
	"strconv"

	"github.com/knightstick/aoc2019/fuelcounterupper"
)

// Part1 find the sum of the fuel requirements for all of the modules on the
// spacecraft
func Part1(moduleMassStrings []string) int {
	if len(moduleMassStrings) == 0 {
		return 0
	}

	moduleMasses := make([]int, 0)

	for _, massString := range moduleMassStrings {
		mass, _ := strconv.Atoi(massString)
		moduleMasses = append(moduleMasses, mass)
	}

	totalFuelRequirement := fuelcounterupper.TotalRequirements(moduleMasses)

	return totalFuelRequirement
}
