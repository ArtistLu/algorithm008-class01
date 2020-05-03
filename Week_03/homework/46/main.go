package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
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
		dfs(append(append([]int{}, nums[0:i]...), nums[i+1:]...), append([]int{nums[i]}, pre...), re)
	}
}
