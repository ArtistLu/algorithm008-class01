package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2, 2, 3, 3}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	re := [][]int{}
	dfs(nums, nil, &re)
	return re
}

func dfs(nums, pre []int, re *[][]int) {
	if len(nums) == 0 {
		*re = append(*re, append([]int{}, pre...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		dfs(append(append([]int{}, nums[0:i]...), nums[i+1:]...), append(pre, nums[i]), re)
	}
}
