package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)
	// Prints O
	var number int8 = 127
	number++
	fmt.Println(number)
	// Prints -128
}
