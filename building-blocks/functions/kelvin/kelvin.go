package main

import "fmt"

// Declares a function that accepts one parameter and returns one result
// kelvinToCelsius converts K to C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// Calls the function passing kelvin as the first argument
func main() {
	// Prints 294° K is 20.850000000000023°C
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "° K is ", celsius, "° C")
}
