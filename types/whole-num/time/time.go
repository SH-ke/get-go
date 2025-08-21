package main

import (
	"fmt"
	"time"
)

func main() {
	// Prints 2370-01-01 00:00:00+0000UTC in the Go Playground
	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}
