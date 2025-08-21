package main

import "fmt"

func main() {
	type celsius float64
	// The underlying type is float64.
	var temperature celsius = 20
	fmt.Println(temperature)
	// Prints 20
}
