package main

import "fmt"

//terraform accomplishes nothing
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
	fmt.Println("in terraform func")
	fmt.Println(planets)
}
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
	terraform(planets)
	fmt.Println("in main: ", planets)
	// Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

	// dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// terraform(dwarfs)
	// Can't use dwarfs (type [5]string]as type [8]string in argument to terraform
}
