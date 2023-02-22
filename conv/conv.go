package conv

func KelvinToCelsius(value float64) float64 {
	return (value + 273.15)
}

func CelsiusToKelvin(value float64) float64 {
	return (value - 273.15)
}

func CelsiusToFarhenheit(value float64) float64 {
	return (value*(9/5) + 32)
}

func FarhenheitToCelsius(value float64) float64 {
	return (value - 32) * (5 / 9)
}

func KelvinToFarhenheit(value float64) float64 {
	return (value-273.15)*(9/5) + 32
}


func FarhenheitToKelvin(value float64) float64 {
	return (value-32)*(5/9) + 273.15
}
