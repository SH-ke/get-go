package main

import "fmt"

func main() {
	// Listing 8.3 Literals of unusual size:constant.go
	// 字面值
	fmt.Println("Andromeda Galaxy is", 2400_0000_0000_0000_0000/29_9792/8_6400, "light days away.")
	// Prints Andromeda Galaxy is 9_2656_8346 light days away.

	// Tisting 8.4 Constants of unusual size:constant.go
	const distance = 2400_0000_0000_0000_0000
	// Prints Andromeda Galaxy is 9_2656_8346 light days away.
	const lightSpeed = 29_9792
	const secondsPerDay = 8_6400
	const days = distance / lightSpeed / secondsPerDay
	fmt.Println("Andromeda Galaxy is", days, "light days away.")
}
