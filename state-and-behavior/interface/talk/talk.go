package main

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {

	var t interface {
		talk() string
	}
	t = martian{}
	// Prints nack nack
	fmt.Println(t.talk())
	t = laser(13)
	// Prints pew pew pew
	fmt.Println(t.talk())
}
