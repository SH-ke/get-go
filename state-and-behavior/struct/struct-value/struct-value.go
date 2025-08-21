package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	bradbury := location{-4.5895, 137.4417}
	curiosity := bradbury
	// Heads east to Yellowknife Bay
	curiosity.long += 0.0106
	// Prints{4.5895 137.4417} {4.5895 137.4523}
	fmt.Println(bradbury, curiosity)

}
