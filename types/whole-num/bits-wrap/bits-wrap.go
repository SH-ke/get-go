package main

import "fmt"

func main() {
	var blue uint8 = 255
	// Prints 11111111
	fmt.Printf("%08b\n", blue)
	blue++
	// Prints 00000000
	fmt.Printf("%08b\n", blue)
}
