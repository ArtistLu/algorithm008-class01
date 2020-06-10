package main

import "fmt"

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, t := range test {
		fmt.Println(t, "->", isPowerOfTwo(t))
	}
}

func isPowerOfTwo(n int) bool {
	re := 0
	for n > 0 {
		n, re = n&(n-1), re+1
		if re > 1 {
			return false
		}
	}

	return re == 1
}
