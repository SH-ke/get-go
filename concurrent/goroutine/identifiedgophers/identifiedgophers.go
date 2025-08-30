package main

import (
	"fmt"
	"time"
)

func main() {
	// 随机顺序启动5个Gopher
	for i := range 5 {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
}
