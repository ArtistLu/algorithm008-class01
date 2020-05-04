package main

func main() {

}

//递归树剪枝
var m = make(map[int]int)

func climbStairs1(n int) int {
	if re, ok := m[n]; ok {
		return re
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	m[n] = climbStairs1(n-1) + climbStairs1(n-2)
	return m[n]
}

func climbStairs2(n int) int {
	if n < 3 {
		return n
	}

	re := 0
	step1, step2 := 1, 2
	for i := 3; i <= n; i++ {
		re = (step1 + step2)
		step1, step2 = step2, re
	}

	return re
}
