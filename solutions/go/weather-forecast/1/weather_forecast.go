// Package weather provides tools to forecast weather.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
    // CurrentLocation represents the current location that we are forecasting weather.
	CurrentLocation  string
)

// Forecast shows the forecasted situation based on city and conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
