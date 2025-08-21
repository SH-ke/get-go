package main

import "fmt"

type kelvin float64

//sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	// To-do:implement a real sensor
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {

	return func() kelvin {
		// Declares and returns an anonymous function
		return s() + offset
	}
}
func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
	// Prints 5
}
