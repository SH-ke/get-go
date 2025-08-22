package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}
func main() {
	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	// Prints &(name:Terry age:16}
	fmt.Printf("%+v\n", terry)
	nathan := person{
		name: "Nathan",
		age:  17}
	nathan.birthday()
	// Prints {name:Nathan age:18)
	fmt.Printf("%+v\n", nathan)
}
