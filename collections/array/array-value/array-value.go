package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	// Copies planets array
	planetsMarkII := planets
	// Makes way for an interstellar bypass
	planets[2] = "whoops"
	fmt.Println(planets)
	// Prints [Mercury Venus whoops Mars Jupiter Saturn Uranus Neptune]
	fmt.Println(planetsMarkII)
	// Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
}
