package main

import "fmt"

func main() {
	test := [][]int{
		{3, 2},
		{7, 3},
	}

	for _, t := range test {
		fmt.Println(t[0], "+", t[1], "=>", uniquePaths(t[0], t[1]))
	}
}

func uniquePaths(m int, n int) int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
	}

	for i := range grid[n-1] {
		grid[n-1][i] = 1
	}

	for i := 0; i < n; i++ {
		grid[i][m-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			grid[i][j] = grid[i+1][j] + grid[i][j+1]
		}
	}

	return grid[0][0]
}
