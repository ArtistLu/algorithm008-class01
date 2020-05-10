package main

import "fmt"

func main() {
	// fmt.Println(generateBoard(7))
	fmt.Println(solveNQueens(7))
}

func solveNQueens(n int) [][]string {
	re := [][]string{}
	bfs(n, 0, 0, 0, 0, generateBoard(n), &re)
	return re
}

func bfs(n, row, colbit, leftbit, rightbit int, board [][]byte, re *[][]string) {
	if n == row {
		*re = append(*re, board2ans(board))
		return
	}

	for i := 0; i < n; i++ {
		bit := 1 << uint(n-i-1)
		if colbit&bit == 0 && leftbit&bit == 0 && rightbit&bit == 0 {
			board[row][i] = 'Q'
			bfs(n, row+1, colbit|bit, (leftbit|bit)<<1, (rightbit|bit)>>1, board, re)
			board[row][i] = '.'
		}
	}
}

func generateBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	return board
}

func board2ans(board [][]byte) []string {
	r := []string{}
	for _, b := range board {
		r = append(r, string(b))
	}
	return r
}
