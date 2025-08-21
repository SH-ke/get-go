package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(29_9792)
	// Prints Andromeda Galaxy is 2400_0000_0000_0000_0000 km away.
	secondsPerDay := big.NewInt(8_6400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	// Prints That is 9_2656_8346 days of travel at light speed.
	fmt.Println("That is", days, "days of travel at light speed.")
	years := new(big.Int)
	daysPerYear := big.NewInt(365)
	years.Div(days, daysPerYear)
	fmt.Println("That is", years, "years of travel at light speed.")
}
