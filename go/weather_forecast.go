// Package weather provides tools to create forecasts.
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string

// CurrentLocation represents current location.
var CurrentLocation string

// Forecast describes current condition at location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
