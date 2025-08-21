package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	// Composite literals are key-value pairs for maps.
	// Prints On average the Earth is 15° C.
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v° C.\n", temp)
	// A little climate change
	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)
	// Prints map[Venus:464 Earth:16 Mars:-65]

	// If you access a key that doesn't exist in the map,the result is the zero value for the type(int):
	moon := temperature["Moon"]
	// Prints 0
	fmt.Println(moon)
	// Go provides the comma,ok syntax,which you can use to distinguish between the "Moon"
	// not existing in the map versus being present in the map with a temperature of 0 C:
	// The comma, ok syntax
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
	// Prints Where is the moon?
}
