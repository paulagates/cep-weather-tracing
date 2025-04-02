package utils

func ConvertCelsiusToFahrenheit(tempC float64) float64 {
	return tempC*1.8 + 32
}

func ConvertCelsiusToKelvin(tempC float64) float64 {
	return tempC + 273.15
}
