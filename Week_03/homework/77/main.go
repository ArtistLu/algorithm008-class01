package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(5, 1))
}

func combine(n int, k int) [][]int {
	re := &[][]int{}
	dfs(1, n, k, []int{}, re)
	return *re
}

func dfs(s, n, k int, current []int, re *[][]int) {
	if len(current) == k {
		*re = append(*re, append([]int{}, current...))
	}

	for i := s; i <= n; i++ {
		dfs(i+1, n, k, append(current, i), re)
	}
}
