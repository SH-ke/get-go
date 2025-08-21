package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
}

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

func (r report) average() celsius {
	return r.temperature.average()
}
func main() {
	t := temperature{high: -1.0, low: -78.0}
	// Prints average -39.5°C
	fmt.Printf("average %v° C\n", t.average())

	// Prints average -39.5°C
	report := report{sol: 15, temperature: t}
	fmt.Printf("average %v C\n", report.temperature.average())
	// 匿名嵌入、具名嵌入
	fmt.Printf("forward average %v C\n", report.average())
}
