package main

import "fmt"

func main() {
	var administrator *string
	scolese := "Christopher J.Scolese"
	administrator = &scolese
	// Prints Christopher J.Scolese
	fmt.Println(*administrator)
	bolden := "Charles F.Bolden"
	administrator = &bolden
	// Prints Charles F.Bolden
	fmt.Println(*administrator)

	// 	Changes to the value of bolden can be made in one place,because the administrator vari-
	// able points to bolden rather than storing a copy:
	bolden = "Charles Frank Bolden Jr."
	// Prints Charles Frank Bolden Jr.
	fmt.Println(*administrator)
	// It's also possible to dereference adninistrator to change the value of bolden indirectly:
	*administrator = "Maj.Gen.Charles Frank Bolden Jr."
	fmt.Println(bolden)
	// Prints Maj.Gen.Charles Frank Bolden Jr.
	// Assigning major to administrator results in a new pointer that's also pointing at the bolden
	// string(see figure 26.3):
	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden)
	// Prints Major General Charles Frank Bolden Jr.

	fmt.Println(administrator == major)
	// Prints true

	lightfoot := "Robert M.Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator == major)
	// Prints false
	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	// Prints Major General Charles Frank Bolden Jr.
	fmt.Println(bolden)
	// Prints Charles Bolden

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)
	// Prints true
	fmt.Println(&charles == &bolden)
	// Prints false
}
