package main

import "fmt"

func main() {
	var bh float64 = 32767
	var h = int16(bh)
	// To-do:add rocket science
	fmt.Println(h)
	h++
	fmt.Println(h)
}
