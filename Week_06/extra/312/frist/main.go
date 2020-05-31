package main

import "fmt"

func main() {
	maxCoins([]int{1, 2, 4, 5, 3, 9})
}

func maxCoinsxxx(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 0
		if i > 0 {
			dp[i-1][i] = 0
		}
		if i > 1 {
			dp[i-2][i] = nums[i-2] * nums[i-1] * nums[i]
		}
	}

	for i := 3; i < n; i++ {
		for j := i - 3; j >= 0; j-- {
			for k := j + 1; k < i; k++ {
				dp[j][i] = max(dp[j][i], dp[j][k]+dp[k][i]+nums[k]*nums[i]*nums[j])
			}
		}
	}
	return dp[0][n-1]
}

func maxCoinsxxxx(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...)
	c := make([][]int, n+2)
	for i := range c {
		c[i] = make([]int, n+2)
	}
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			for k := i; k <= j; k++ {
				c[i][j] = max(c[i][j], c[i][k-1]+nums[i-1]*nums[k]*nums[j+1]+c[k+1][j])
			}
		}
	}
	return c[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...)
	c := make([][]int, n+2)
	for i := range c {
		c[i] = make([]int, n+2)
	}
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			fmt.Println("[", i, j, "]")
			for k := i; k <= j; k++ {
				c[i][j] = max(c[i][j], c[i][k-1]+nums[i-1]*nums[k]*nums[j-1]+c[k+1][j])
			}
		}
	}
	fmt.Println(c)
	return c[1][n]
}
