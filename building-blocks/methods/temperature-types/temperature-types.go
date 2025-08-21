package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius convertsKtoC
func kelvinToCelsius(k kelvin) celsius {
	// A type conversion is necessary
	return celsius(k - 273.15)
}
func main() {
	// The argument must be of type kelvin.
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, "° K is ", c, "° C")
	// Prints 294° K is 20.850000000000023° C
}
