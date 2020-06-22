package main

import "fmt"

func main() {
	test := [][]string{
		{"rabbbit", "rabbit"},
	}

	for _, t := range test {
		fmt.Println(numDistinct(t[0], t[1]))
	}

}

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for r := range dp {
		dp[r] = make([]int, len(s)+1)
	}

	for j := 0; j <= len(s); j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(t)][len(s)]
}
