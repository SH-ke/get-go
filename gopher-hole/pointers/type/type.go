package main

import "fmt"

func main() {
	answer := 42
	address := &answer
	// Prints address is a *int
	fmt.Printf("address is a %T\n", address)
}
