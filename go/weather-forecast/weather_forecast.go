// Package weather provides forecast of weather.
package weather

var (
	// CurrentCondition store the information about current weather condition.
	CurrentCondition string
	// CurrentLocation stores the information about location where user currently at.
	CurrentLocation string
)

// Forecast take 2 argurment city and condition as string and return weather condition as string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
