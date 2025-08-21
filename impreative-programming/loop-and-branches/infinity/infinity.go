package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var degrees = 0
	round := 1
	for {
		fmt.Printf("Degrees: %v, Round: %v\n", degrees, round)
		degrees++
		if degrees > 360 {
			degrees = 0
			round++
			if rand.Intn(2) == 0 {
				break
			}
		}
	}

}
