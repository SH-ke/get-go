package main

import "fmt"

// Assigns an anonymous function to a variable
var f = func() {
	fmt.Println("Dress up for the masquerade.")
}

// Prints Dress up for the masquerade.
func main() {
	f()
}
