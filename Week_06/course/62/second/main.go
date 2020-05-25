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
	status := make([]int, m)
	status[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			status[j] += status[j-1]
		}
	}

	return status[m-1]
}
