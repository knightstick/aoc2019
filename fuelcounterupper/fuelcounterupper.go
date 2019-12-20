package fuelcounterupper

func TotalRequirements(moduleMasses []int) (total int) {
	for _, mass := range moduleMasses {
		total += FuelRequirement(mass)
	}

	return
}

func FuelRequirement(mass int) int {
	return (mass / 3) - 2
}
