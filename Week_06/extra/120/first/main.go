package main

import "fmt"

func main() {
	re := []int{1, 2, 3}
	s := append(re, 9)
	t := append(s, 9)
	fmt.Println(re, s, t)
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
