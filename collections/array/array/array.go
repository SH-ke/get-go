package main

import "fmt"

func main() {
	var planets [8]string
	// Assigns a planet at index O
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	// Retrieves the planet at index 2
	earth := planets[2]
	fmt.Println(earth)
	// Prints Earth

	fmt.Println(len(planets))     // Prints 8
	fmt.Println(planets[3] == "") // Prints true
}
