package main

import "fmt"

func main() {
	type celsius float64
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)

	var warmUp float64 = 10
	// Invalid operation: mismatched types
	// temperature += warmUp

	// var warmUp float64 = 10
	temperature += celsius(warmUp)
	fmt.Println(temperature)
}
