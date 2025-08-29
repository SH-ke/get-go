package main

import "fmt"

func main() {
	var v interface{}
	// Prints <nil><nil>true
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	// Prints *int <nil> false

	fmt.Printf("%#v\n", v)
	// Prints(*int)(nil)
}
