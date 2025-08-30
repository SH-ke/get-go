package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// Makes the channel to communicate over
	for i := range 5 {
		go sleepyGopher(i, c)
	}
	for range 5 {
		// Receives a value from a channel
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}
func sleepyGopher(id int, c chan int) {
	// Declares the channel as an argument
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id
	// Sends a value back to main
}
