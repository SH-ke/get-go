package main

import (
	"fmt"
	"math"
	"sort"
)

type world struct {
	radius float64
}

type location struct {
	lat, long float64
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	// Uses the world's radius field
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func main() {
	var mars = world{radius: 3389.5}

	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	// Uses the distance method on mars
	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2fkm\n", dist)
	// Prints 9669.71 km

	var earth = world{radius: 6371}
	wuhan := location{30.5331, 114.3116}
	beijing := location{39.8968, 116.3260}
	me := location{30.7375, 115.66722222222222}
	// Uses the distance method on earth
	// locations := []location{wuhan, beijing, me}
	location_map := map[string]location{"wuhan": wuhan, "beijing": beijing, "me": me}

	labels := make([]string, 0)
	for name := range location_map {
		labels = append(labels, name)
	}
	fmt.Println(labels)
	sort.Strings(labels)
	n := len(labels)
	for i := range n {
		a := labels[i]
		for j := i + 1; j < n; j++ {
			b := labels[j]
			dist := earth.distance(location_map[a], location_map[b])
			fmt.Printf("the distance between %v and %v is: %.2fkm\n", a, b, dist)
		}
	}
}
