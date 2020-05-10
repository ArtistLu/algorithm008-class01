package main

func main() {

}

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		ugly[i] = minNum(ugly[p2]*2, minNum(ugly[p3]*3, ugly[p5]*5))
		if ugly[i] == ugly[p2]*2 {
			p2++
		}
		if ugly[i] == ugly[p3]*3 {
			p3++
		}
		if ugly[i] == ugly[p5]*5 {
			p5++
		}
	}
	return ugly[n-1]
}

func minNum(i, j int) int {
	if i > j {
		return j
	}
	return i
}
