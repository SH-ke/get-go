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

//newLocation from latitude,longitude d/m/s coordinates.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {

	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)
	me := newLocation(coordinate{30, 44, 15, 'N'}, coordinate{115, 40, 2, 'E'})
	fmt.Println(me)
	// Prints (-4.5895 137.4417)
}
