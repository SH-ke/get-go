package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	p.age++
	// nil pointer dereference
}
func main() {

	var nobody *person
	fmt.Println(nobody)
	// Prints <nil
	nobody.birthday()
}
