// Package weather provides a function that returns a string containing the current location and current weather condition.
package weather

// CurrentCondition represents the current weather condition at a specified location.
var CurrentCondition string

// CurrentLocation represents the name of the location for which the current weather condition is being reported.
var CurrentLocation string

// Forecast returns a string that reports the current location (city) and the current weather condition for that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
