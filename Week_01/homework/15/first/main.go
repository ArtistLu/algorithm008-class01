package main

import (
	"fmt"
	"sort"
)

func main() {
	test := [][]int{
		{1, -2, 1, 2, -4, 2, 9},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, -1, -1, -1, 2, 2, 2, -2, -2, -2},
	}

	for _, t := range test {
		fmt.Println(threeSum(t))
	}
}

//https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i, t := range nums {
		l, r := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			sum := nums[l] + nums[r] + t
			if sum == 0 {
				ans = append(ans, []int{t, nums[l], nums[r]})
				r--
				l++
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}
				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return ans
}
