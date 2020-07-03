package kmp

//参考 https://mp.weixin.qq.com/s/r9pbkMyFyMAvmkf4QnL-1g
type KMP struct {
	p  string
	dp [][]int
}

func NewKmp(p string) *KMP {
	return &KMP{
		p:  p,
		dp: kmp(p),
	}
}

func kmp(p string) [][]int {
	dp := make([][]int, len(p))
	for i := range dp {
		dp[i] = make([]int, 256)
	}
	dp[0][p[0]] = 1
	x := 0
	for i := 1; i < len(p); i++ {
		for j := 0; j < 256; j++ {
			dp[i][j] = dp[x][j]
		}
		dp[i][p[i]] = i + 1
		x = dp[x][p[i]]
	}
	return dp
}

func (k *KMP) Search(s string) int {
	j := 0
	for i := 0; i < len(s); i++ {

		j = k.dp[j][s[i]]
		if j == len(k.p) {
			return i - len(k.p) + 1
		}
	}
	return -1
}
