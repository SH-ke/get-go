// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	const lightSpeed = 29_9792 //km/s
	var distance = 5600_0000   //km
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 4_0100_0000
	fmt.Println(distance/lightSpeed, "seconds")
}
