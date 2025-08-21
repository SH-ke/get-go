package main

import "fmt"

func main() {
	// If command is equal to “go east'"
	var command = "go east"
	if command == "go east" {

		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		// Otherwise,if command is equal to“go inside"
		fmt.Println("You enter the cave where you live out the rest of your1ife.")
	} else {
		// Or,if anything else
		fmt.Println("Didn't quite get that.")
	}
}
