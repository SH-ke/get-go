package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	fmt.Printf("%+v, %[1]v\n", timmy)
	// Prints &(name:Timothy
	// superpower:flying
	// age:10}
}
