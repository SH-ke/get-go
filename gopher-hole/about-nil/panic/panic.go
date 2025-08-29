package main

import "fmt"

func main() {
	var nowhere *int
	// Prints <nil>
	fmt.Println(nowhere)
	// Panic:nil pointer dereference
	fmt.Println(*nowhere)

}
