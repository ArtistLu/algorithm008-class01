package main

import "fmt"

func main() {
	testMatrix := [][]int{
		{9, 9, 9},
		{1, 2, 3},
		{},
	}

	for _, v := range testMatrix {
		fmt.Print(v, " -> ")
		fmt.Println(plusOne(v))
	}
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
