package main

import (
	"fmt"
	"math/rand"
)

// era is available throughout the package.
var era = "AD"

func main() {
	// era and year are in scope.
	year := 2018

	// era,year,and month are in scope.
	switch month := rand.Intn(12) + 1; month {

	case 2:
		// era,year,month,and day are in scope.
		day := rand.Intn(28) + 1

		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		// It's a new day.
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
	// month and day are out of scope.
}
