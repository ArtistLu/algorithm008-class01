package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{1, 2, 4, 5, 3, 9}))
}

func maxCoins(nums []int) int {
	//dp[i][j] 表示下标i到j开区间戳破气球的最大硬币值
	//dp[i][j]  i < k < j k表示最后戳破的气球
	//dp[i][j] = max(dp[i][j], dp[i][k] + nums[i] * nums[k] * nums[j] + dp[k][j])
	//dp[0][n + 1] 要求的结果

	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...)

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+nums[i]*nums[k]*nums[j]+dp[k][j])
			}
		}
	}

	return dp[0][n+1]
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
