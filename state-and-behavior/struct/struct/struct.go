package main

import "fmt"

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}
	// Assigns values to fields of the structure
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	// Prints -4.5895137.4417
	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
	// Prints {-4.5895 137.4417}
}
