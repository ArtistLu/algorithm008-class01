package main

import (
	"fmt"
	"math"
)

func main() {

}

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	bfs(n, 0, 0, 0, 0, generateBoard(n), &ans)
	return ans
}

func bfs(n, row, colbit, leftbit, rightbit int, board []string, ans *[][]string) {
	if row == n {
		*ans = append(*ans, board)
		return
	}

	for i := 0; i < n; i++ {
		bit := 1 << (n - i - 1)
		if bit&colbit == 0 && bit&leftbit == 0 && bit&rightbit == 0 {
			nextBoard := append([]string{}, board...)
			nextBoard[row] = board[row][0:i] + "Q" + board[row][i+1:]
			bfs(n, row+1, colbit|bit, (leftbit|bit)<<1, (rightbit|bit)>>1, nextBoard, ans)
		}
	}
}

func generateBoard(n int) []string {
	var row string
	for i := 0; i < n; i++ {
		row += "."
	}
	board := []string{}
	for i := 0; i < n; i++ {
		board = append(board, row)
	}
	return board
}
