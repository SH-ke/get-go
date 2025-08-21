package main

import "fmt"

func main() {
	var green uint8 = 3
	// Prints 00000011
	fmt.Printf("%08b\n", green)
	green++
	// Prints 00000100
	fmt.Printf("%08b\n", green)
}
