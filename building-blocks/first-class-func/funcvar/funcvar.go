package main

import "fmt"

func main() {
	// Assigns an anonymous function to a variable
	f := func(message string) {
		fmt.Println(message)
	}

	// Prints Go to the party.
	f("Go to the party.")
}
