package day1

import (
	"strconv"

	"github.com/knightstick/aoc2019/fuelcounterupper"
)

// Part1 find the sum of the fuel requirements for all of the modules on the
// spacecraft
func Part1(moduleMassStrings []string) int {
	return fuelcounterupper.TotalRequirementsWithoutFuelNeedingFuel(moduleMassesFromStrings(moduleMassStrings))
}

// Part2 finds the sum of the fuel requirements for all of the modules on the
// spacecraft, including the need of the weight from the fuel needing more fuel
func Part2(moduleMassStrings []string) int {
	return fuelcounterupper.TotalRequirements(moduleMassesFromStrings(moduleMassStrings))
}

func moduleMassesFromStrings(moduleMassStrings []string) (moduleMasses []int) {
	if len(moduleMassStrings) == 0 {
		return []int{}
	}

	for _, massString := range moduleMassStrings {
		mass, _ := strconv.Atoi(massString)
		moduleMasses = append(moduleMasses, mass)
	}

	return
}
