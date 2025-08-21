package main

import (
	"fmt"
	"strings"
)

type starship struct {
	laser
}
type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	s := starship{laser(3)}
	// Prints pew pew pew
	fmt.Println(s.talk())
	shout(s)
	// Prints PEW PEW PEW
}
