package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

// stardate returns a fictional measure of time.
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	// There are 668 sols in a Martian year.
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
	// The hour is unknown.
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))
	// Prints 1219.2 Curiosity has landed
	s := sol(1422)

	fmt.Printf("%.1f Happy birthday\n", stardate(s))
	// Prints 1086.0 Happy birthday
}
