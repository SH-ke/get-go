package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	// Makes a new slice rather than modifying worlds directly
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
func main() {
	// The worlds parameter is a slice of strings that contains zero or more arguments passed to terraform
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)
	// Prints [New Venus New Mars]
	// To pass a slice instead of multiple arguments,expand the slice with an ellipsis:
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
	// Prints [New Venus New Mars New Jupiter]
}
