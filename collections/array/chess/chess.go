package main

import "fmt"

func main() {
	var board [8][8]string
	// An array of eight arrays of eight strings
	board[0][0] = "r"
	board[0][7] = "r"
	// Places a rook at a [row][column]coordinate
	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Print(board)
}
