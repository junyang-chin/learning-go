// Package weather provides weather forecasts of a city.
package weather

// CurrentCondition represents the weather condition.
var CurrentCondition string

// CurrentLocation respresents the name place on a map.
var CurrentLocation string

// Forecast takes a name of a place and it's weather condition and returns a readable message.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
