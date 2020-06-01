package main

import "fmt"

func main() {
	test := [][]int{
		{2, 3, 0, -5},
		{8, 3, 0, -5},
		{1, -1, 2, -2, 4, -4},
	}

	for _, t := range test {
		fmt.Println(t, maxSubArray(t))
	}
}

func maxSubArray(nums []int) int {
	//dp[i] 包含nums[i]子数组的最大值
	//dp[i] = nums[i] + (dp[i-1] > 0 ? dp[i-1] : 0)

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	re := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + max(dp[i-1], 0)
		re = max(re, dp[i])
	}
	return re
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
