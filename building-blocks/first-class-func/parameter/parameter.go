package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	// measureTemperature accepts a function as the second parameter.
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v K\n", k)
		time.Sleep(time.Second)
	}
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	// Passes the name of a function to a function
	measureTemperature(3, fakeSensor)
}
