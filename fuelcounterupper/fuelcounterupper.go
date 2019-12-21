package fuelcounterupper

func TotalRequirements(moduleMasses []int) (total int) {
	for _, mass := range moduleMasses {
		total += FuelRequirement(mass)
	}

	return
}

func TotalRequirementsWithoutFuelNeedingFuel(moduleMasses []int) (total int) {
	for _, mass := range moduleMasses {
		total += FuelRequirementWithoutFuelNeedingFuel(mass)
	}

	return
}

func FuelRequirement(mass int) int {
	return calculateFuelForMass(mass, []int{})
}

func calculateFuelForMass(mass int, components []int) int {
	component := FuelForMass(mass)
	components = append(components, component)

	if component < 1 {
		return sum(components)
	}

	return calculateFuelForMass(component, components)
}

func sum(list []int) (result int) {
	for _, i := range list {
		result += i
	}
	return
}

func FuelRequirementWithoutFuelNeedingFuel(mass int) int {
	return FuelForMass(mass)
}

func FuelForMass(mass int) int {
	fuel := (mass / 3) - 2

	if fuel < 0 {
		return 0
	}

	return fuel
}
