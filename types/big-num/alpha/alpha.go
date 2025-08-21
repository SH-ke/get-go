package main

import "fmt"

func main() {
	const lightSpeed = 29_9792 //km/s
	// Prints Alpha Centauri is 41_3000_0000_0000 km away
	const secondsPerDay = 8_6400
	var distance int64 = 41.3e12
	// Prints That is 1594 days of travel at light speed
	fmt.Println("Alpha Centauri is", distance, "km away.")
	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
	fmt.Println("That is", days/365, "years of travel at light speed.")
}
