package main

import (
	"errors"
	"fmt"
	"strings"
)

type SudokuError []error

const rows, columns = 9, 9

// Grid is a Sudoku grid
type Grid [rows][columns]int8

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

// Error returns one or more errors separated by commas.
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
		// Converts the errors to strings
	}
	return strings.Join(s, ",")
}
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}
func (g *Grid) Set(row, column int, digit int8) error {
	// Returns type is error
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
	// Returns nil
}

func main() {
	var g Grid
	err := g.Set(10, 0, 15)
	if err != nil {
		// err.(SudokuError) 是类型断言（Type Assertion）的语法，而 ok 是用于判断类型断言是否成功的标志
		// 检查一个接口值 err 是否持有特定类型 SudokuError 的值
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
}
