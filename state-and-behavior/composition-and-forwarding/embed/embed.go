package main

import "fmt"

type report struct {
	sol int
	temperature
	location
}

// A temperature type embedded into report

// The temperature field is a structure of type temperature.
type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	report := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0}}
	// Prints average -39.5° C
	fmt.Printf("average %v° C\n", report.average())
	// 匿名嵌入、具名嵌入
	fmt.Printf("average %v° C\n", report.temperature.average())
	// Prints average 39.5° C

	fmt.Printf("%v° C\n", report.high)
	// Prints -1° C
	report.high = 32
	fmt.Printf("%v° C\n", report.temperature.high)
	// Prints 32° C
}
