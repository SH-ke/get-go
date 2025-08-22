package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)
	// Prints 0x1040c108
	// 0xc00000a098 random
	address := &answer
	fmt.Println(*address)
	// Prints 42
}
