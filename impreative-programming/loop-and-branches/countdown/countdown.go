package main

import (
	"fmt"
	"time"
)

func main() {
	// Declares and initializes
	var count int = 10
	// A condition
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
		// Decrements count;
		// otherwise it will
		// loop forever
	}
	fmt.Println("Liftoff!")
}
