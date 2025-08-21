package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// Length 4,capacity 4
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
	// Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

	// Length 4,capacity 8
	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
	// Prints [Mercury Venus Earth Mars Ceres Saturn Uranus Neptune]
}
