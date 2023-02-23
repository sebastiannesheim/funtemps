package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
)

var fahr float64
var kel float64
var cel float64
var out string
var funfacts string

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")

}

func main() {

	flag.Parse()

	if isFlagPassed("funfacts") {
		funfactsAbout := GetFunFacts(funfacts)
		for i := 0; i < len(funfactsAbout); i++ {
			fmt.Println(funfactsAbout[i])
		}
	}

	// kelvin to celsius
	if isFlagPassed("K") && out == "C" {
		output := conv.KelvinToCelsius(kel)
		fmt.Printf("%v K er %v C", kel, output)
	}

	// celsius to kelvin
	if isFlagPassed("C") && out == "K" {
		output := conv.CelsiusToKelvin(cel)
		fmt.Printf("%v C er %v K", cel, output)
	}

	// celsius to fahrenheit
	if isFlagPassed("C") && out == "F" {
		output := conv.CelsiusToFahrenheit(cel)
		fmt.Printf("%v C er %v F", cel, output)
	}

	// fahrenheit to celsius
	if isFlagPassed("F") && out == "C" {
		output := conv.FahrenheitToCelsius(fahr)
		fmt.Printf("%v F er %v C", fahr, output)
	}

	// kelvin to fahrenheit
	if isFlagPassed("K") && out == "F" {
		output := conv.KelvinToFahrenheit(kel)
		fmt.Printf("%v K er %v F", kel, output)
	}

	// fahrenheit to kelvin
	if isFlagPassed("F") && out == "K" {
		output := conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%v F er %v K", fahr, output)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
