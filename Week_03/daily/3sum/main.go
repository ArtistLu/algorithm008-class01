package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/3sum/
func main() {
	test := [][]int{
		{1, 0, -1, 2, 1, -1, -1},
		{},
		{1, 2},
		{1, 2, 9},
		{-9, 4, 3, 5, -4, 1},
		{-1, 0, 1, 2, -1, -4},
		{-2, 0, 1, 1, 2},
	}

	for _, t := range test {
		fmt.Print(t, "->")
		re := threeSum(t)
		for _, r := range re {
			fmt.Print(r, "| ")
		}
		fmt.Println()
	}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	re := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				re = append(re, []int{nums[i], nums[l], nums[r]})
				l, r = l+1, r-1
				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}

			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}

	}

	return re
}
