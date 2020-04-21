### 每日一题
- [day2-1](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)
- [day2-2](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)
- [day2-3](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)
  
```
  func getKthMagicNumber(k int) int {
	pt, pf, ps := 0, 0, 0
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i < k; i++ {
		dp[i] = minNum(dp[pt]*3, minNum(dp[pf]*5, dp[ps]*7))
		if dp[i] == dp[pt]*3 {
			pt++
		}
		if dp[i] == dp[pf]*5 {
			pf++
		}
		if dp[i] == dp[ps]*7 {
			ps++
		}
	}

	return dp[len(dp)-1]
}

func minNum(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

```