/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Slice 包含任何 Type（包括其它 Slice）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[0][2] = "X"
	board[1][0] = "O"
	board[1][2] = "X"
	board[2][2] = "O"

	for i := 0; i < cap(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}

	fmt.Println(board)
}
