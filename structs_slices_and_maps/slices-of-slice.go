package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[0][1] = "X"
	board[0][2] = "O"
	board[1][0] = "X"

	for _, v := range board {
		fmt.Println(strings.Join(v, ""))
	}
}
