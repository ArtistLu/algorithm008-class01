package main

func main() {

}

//https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/submissions/
func lastRemaining(n int, m int) int {
	p := 0
	for i := 1; i <= n; i++ {
		p = (p + m) % i
	}

	return p
}
