package main

import "fmt"

func main() {
	yesNo := "no"
	launch := (yesNo == "yes")
	// Prints Ready for launch: false
	fmt.Println("Ready for launch:", launch)
}
