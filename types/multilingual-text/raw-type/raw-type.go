package main

import "fmt"

func main() {
	fmt.Printf("%v is a %[1]T\n", "literal string")
	// Prints literal string is a string
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
	// Prints raw string literal is a string
}
