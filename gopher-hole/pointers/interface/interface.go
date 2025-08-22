package main

import (
	"fmt"
	"strings"
)

type laser int
type talker interface {
	talk() string
}

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))

}
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	pew := laser(2)
	// Prints PEW PEW
	shout(&pew)
}
