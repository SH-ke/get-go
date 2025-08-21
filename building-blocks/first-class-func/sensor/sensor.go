package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func realSensor() kelvin {
	// To-do:implement a real sensor
	return 0
}

func main() {
	sensor := fakeSensor
	fmt.Println(sensor())
	// Assigns a function to a variable
	sensor = realSensor
	fmt.Println(sensor())
}
