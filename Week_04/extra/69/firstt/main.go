package main

import "fmt"

func main() {
	test := []int{0, 1, 2, 4, 8, 9, 11}
	for _, x := range test {
		fmt.Println(x, "->", mySqrt(x))
	}
}

func mySqrt(x int) int {
	s, e := 0, x>>1+1
	for {
		m := s + (e-s)>>1
		if m*m > x {
			e = m - 1
		} else if (m+1)*(m+1) > x {
			return m
		} else {
			s = m + 1
		}
	}
}
