package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity) // Prints{-1.9462 354.4734}
	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)
	// Prints {4.5 135.9}

	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)
	// Prints {-14.5684 175.472636}
}
