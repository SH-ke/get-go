package main

import "fmt"

func main() {
	// switch statement
	room := "lake"

	switch room {
	case "cave":
		fmt.Println("You find yourself in a dimly lit cave.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}
