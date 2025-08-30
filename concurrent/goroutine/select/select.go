package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}
func main() {
	c := make(chan int)
	for i := range 5 {
		go sleepyGopher(i, c)
	}
	// time.After 函数会返回一个 channel，该 channel 会在指定的时间后发送一个当前时间的时间戳
	timeout := time.After(2 * time.Second)
	for range 5 {
		select {
		// The select statement
		case gopherID := <-c:
			// Waits for a gopher to wake up
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case timeoutInfo := <-timeout:
			// Waits for time to run out
			fmt.Println("my patience ran out", timeoutInfo)
			return
			// Gives up and returns
		}
	}
}
