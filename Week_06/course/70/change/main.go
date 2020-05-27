package main

import "fmt"

func main() {
	test := [][]int{
		{1, 3},
		{2, 3},
		{3, 3},
		{4, 3},
		{5, 3},
		{6, 3},
	}

	for _, t := range test {
		fmt.Println(t[0], "+", t[1], climbStairsChange(t[0], t[1]))
	}
}

//参数n代表楼梯阶数， s代表不同的走法。s=3代表有每次走1、2或3阶
func climbStairsChange(n, s int) int {
	dp := make([][]int, s)
	for i := range dp {
		dp[i] = make([]int, n+s)
	}

	dp[0][0], dp[1][1], dp[2][2] = 1, 1, 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= s; j++ {
			if dp[j-1][i-1] == 0 {
				continue
			}
			for k := 1; k <= s; k++ {
				if k == j || i+k > n {
					continue
				}
				dp[k-1][i+k-1]++
			}
		}
	}

	ans := 0
	for i := 0; i < s; i++ {
		ans += dp[i][n-1]
	}

	return ans
}

//s=3 dp数组初始值
//   0  1  2  3  4  5  6  7  8
// -------------------------------------------------
// 0  1
// -------------------------------------------------
// 1     1
// -------------------------------------------------
// 2        1
// --------------------------------------------------
