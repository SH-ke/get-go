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

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	// Prints
	// {
	// 	sol:15
	// 	temperature:{high:-1 low:-78}
	// 	location:{lat:-4.5895 long:137.4417}
	// }
	fmt.Printf("a balmy %v C\n", report.temperature.high)
	// Prints a balmy -1°C
}
