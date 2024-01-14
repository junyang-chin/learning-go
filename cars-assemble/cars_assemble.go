package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(producedPerHour int, successRate float64) float64 {

	return float64(producedPerHour) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(producedPerHour int, successRate float64) int {
	producedPerMinute := float64(producedPerHour) / 60

	return int(producedPerMinute * successRate / 100)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTens := carsCount / 10
	remainder := carsCount % 10

	return uint(groupsOfTens)*95000 + uint(remainder)*10000
}
