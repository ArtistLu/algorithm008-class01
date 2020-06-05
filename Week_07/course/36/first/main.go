package main

import "fmt"

func main() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	n := [][]int{}

	copy(n, m)
	fmt.Println(n, m)
}

func isValidSudoku(board [][]byte) bool {
	m := make([][]int, 9)
	for i := range m {
		m[i] = make([]int, 3)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			n := int(board[i][j] - '0')
			cur := m[n-1]
			rowbit, colbit, arebit := 1<<(9-i-1), 1<<(9-j-1), 1<<(9-(i/3*3+j/3)-1)
			if cur[0]&rowbit != 0 || cur[1]&colbit != 0 || cur[2]&arebit != 0 {
				return false
			}
			cur[0], cur[1], cur[2] = cur[0]|rowbit, cur[1]|colbit, cur[2]|arebit
		}
	}

	return true
}
