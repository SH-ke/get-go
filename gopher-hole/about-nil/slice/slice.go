package main

import "fmt"

func main() {
	var soup []string
	// Prints true
	fmt.Println(soup == nil)
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	// Prints 0
	fmt.Println(len(soup))
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)
	// Prints [onion carrot celery]
}
