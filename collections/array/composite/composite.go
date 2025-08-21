package main

import "fmt"

func main() {
	planets := [...]string{ // The Go compiler counts the elements.
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune", // The trailing comma is required.
	}
	fmt.Print(planets)

}
