package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting the program")
	// The goroutine is started.
	go sleepyGopher()
	// Waiting for the gopher to snore
	time.Sleep(4 * time.Second)
	// When we get here, all the goroutines are stopped.
	fmt.Println("Ending the program")
	fmt.Println("Starting 5 gophers")
	for range 5 {
		go sleepyGopher()
	}
	// time.Sleep(3 * time.Second)
	time.Sleep(4 * time.Second)
}
func sleepyGopher() {
	time.Sleep(3 * time.Second)
	// The gopher sleeps.
	fmt.Println("... snore ...")
}
