package main

import "fmt"

//coordinate in degrees,minutes,seconds in a N/5/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

//decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {

	//Bradbury Landing:435'22.2"S,13726'30.1"E
	lat := coordinate{4, 35, 22.2, 'S'}

	long := coordinate{137, 26, 30.12, 'E'}
	// Prints -4.5895 137.4417
	fmt.Println(lat.decimal(), long.decimal())

	curiosity := location{lat.decimal(), long.decimal()}
	fmt.Println(curiosity)
}
