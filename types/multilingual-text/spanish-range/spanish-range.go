package main

import "fmt"

func main() {
	question := "¿Cómo estás?"
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}

	fmt.Println("some rune for 2 bytes")
	// some rune for 2 bytes
	runes := []rune(question)
	for i, r := range runes {
		fmt.Printf("%v %c %d\n", i, r, len([]byte(string(r))))
	}
}
