package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := make(map[float64][]float64)
	// A map with float64 keys and []float64 values
	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	// Rounds temperatures down to -20,-30,and so on
	for g, temperatures := range groups {
		fmt.Printf("%v:%v\n", g, temperatures)
	}
}
