package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {
	// This argument is a slice,not an array.
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

// Planets surrounded by space
func main() {

	planets := []string{"Venus ", "Earth ", "Mars"}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
	// Prints VenusEarthMars
}
