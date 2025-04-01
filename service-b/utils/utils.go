package utils

// ConvertCelsiusToFahrenheit converts a temperature in Celsius to Fahrenheit.
func ConvertCelsiusToFahrenheit(tempC float64) float64 {
	return tempC*1.8 + 32
}

// ConvertCelsiusToKelvin converts a temperature in Celsius to Kelvin.
func ConvertCelsiusToKelvin(tempC float64) float64 {
	return tempC + 273.15
}
