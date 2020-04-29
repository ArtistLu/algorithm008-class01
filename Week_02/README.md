### 每日一题
- [day2-1](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)
- [day2-2](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)
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
- [day2-3](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)
- [day3-1](https://leetcode-cn.com/problems/remove-outermost-parentheses/)
- [day3-2](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)
- day4-1 412 [Fizz Buzz](https://leetcode-cn.com/problems/fizz-buzz/)
- [day4-2](https://leetcode-cn.com/problems/add-digits)
- [day4-3](https://leetcode-cn.com/problemsmove-zeroes)
- [day5-1](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
- [day5-2](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)
- [day5-3](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/)
- [day5-4](https://leetcode-cn.com/problems/balanced-binary-tree/)
- [day6-1](https://leetcode-cn.com/problems/zero-matrix-lcci/)
- [day6-2](https://leetcode-cn.com/problems/minimum-absolute-difference/)
- [day6-3](https://leetcode-cn.com/problems/diameter-of-binary-tree/)
- [day7-1](https://leetcode-cn.com/problems/sort-colors/)
- [day7-2](https://leetcode-cn.com/problems/merge-k-sorted-lists/)
