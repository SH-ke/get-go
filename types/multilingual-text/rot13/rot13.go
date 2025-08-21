package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	// Iterates through each ASCll character
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			// Leaves spaces and punctuation as they are
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
