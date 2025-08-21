package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "whoops"
	// Prints map[Earth:whoops Mars:Sector ZZ9]
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	delete(planets, "Earth")
	// Removes Earth from the map
	fmt.Println(planetsMarkII)
	// Prints map[Mars:Sector ZZ9]-
}
