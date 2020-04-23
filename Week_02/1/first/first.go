package main

import (
	"fmt"
)

func main() {
	testMatrix := [][]int{
		{1, 3, 4, 9, 9},
		{1, 2, 4, 7, 9},
		{1, 3, 4, 9, 9},
	}

	test := []int{
		7, 6, 12,
	}

	for i, v := range testMatrix {
		idx := twoSum(v, test[i])
		fmt.Println(v, " | ", test[i], "->", []int{v[idx[0]], v[idx[1]]})
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}

	return []int{}
}
