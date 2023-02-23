package conv

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func CelsiusToFahrenheit(value float64) float64 {
	return value*9/5 + 32
}

func FahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

func KelvinToFahrenheit(value float64) float64 {
	return (value-273.15)*9/5 + 32
}


func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*5/9 + 273.15
}
