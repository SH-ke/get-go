// Listing 10.5 Converting a Boolean to a string:launch.go
package main

import "fmt"

func main() {
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)
	// Prints Ready for launch:false
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	// Prints Ready for launch:no
	fmt.Println("Ready for launch:", yesNo)
}
