package main

import (
	"fmt"
	"testing"

	"github.com/knightstick/aoc2019/fuelcounterupper"
)

func TestFuelCounterUpperTotalRequirements(t *testing.T) {
	masses := []int{12, 14, 1969}
	actual := fuelcounterupper.TotalRequirements(masses)
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
			actual := fuelcounterupper.FuelRequirement(tt.mass)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}
