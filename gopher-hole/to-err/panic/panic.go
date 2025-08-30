package main

import "fmt"

func main() {
	stop := func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			// Recovers from panic
		}
	}

	defer stop()
	// stop()
	// Prints I forgot my towel
	panic("I forgot my towel")
	// Causes panic
	fmt.Println("This won't be printed")
}
