package main

import "fmt"

func main() {
	canada := "Canada"
	var home *string
	// Prints home is a *string
	fmt.Printf("home is a %T, value: %[1]v\n", home)

	home = &canada
	// Prints Canada
	fmt.Println(*home)
}
