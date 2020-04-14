package main

import "fmt"

func main() {
	testM := [][]int{
		{1, 2},
		{1},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4},
	}

	testK := []int{1, 2, 3, 4, 5}

	for _, m := range testM {
		for _, k := range testK {
			var s = m
			fmt.Print(s, " right move ", k)
			rotate(s, k)
			fmt.Print("->", s)
			fmt.Println()
		}
	}

}

func rotate(nums []int, k int) {
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
}
