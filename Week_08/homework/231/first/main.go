package main

import "fmt"

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, t := range test {
		fmt.Println(t, "->", isPowerOfTwo(t))
	}
}

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0 && n != 0
}
