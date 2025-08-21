package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// Sorts planets alphabetically
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
	// Prints [Earth Jupiter Mars Mercury Neptune Saturn Uranus Venus]
}
