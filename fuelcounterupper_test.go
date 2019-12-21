package main

import (
	"fmt"
	"testing"

	"github.com/knightstick/aoc2019/fuelcounterupper"
)

func TestFuelCounterUpperTotalRequirements(t *testing.T) {
	masses := []int{12, 14, 1969}
	actual := fuelcounterupper.TotalRequirementsWithoutFuelNeedingFuel(masses)
	expected := 2 + 2 + 654

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestFuelCounterUpperFuelRequirement(t *testing.T) {
	var reqTests = []struct {
		mass     int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tt := range reqTests {
		t.Run(fmt.Sprintf("mass of %d", tt.mass), func(t *testing.T) {
			actual := fuelcounterupper.FuelRequirementWithoutFuelNeedingFuel(tt.mass)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

func TestFuelCounterUpperFuelRequirementWithFuelNeedingFuel(t *testing.T) {
	var reqTests = []struct {
		mass     int
		expected int
	}{
		{14, 2},
		{1969, 654 + 216 + 70 + 21 + 5},
		{100756, 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2},
	}

	for _, tt := range reqTests {
		t.Run(fmt.Sprintf("mass of %d", tt.mass), func(t *testing.T) {
			actual := fuelcounterupper.FuelRequirement(tt.mass)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}
