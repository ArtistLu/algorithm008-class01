package main

import "fmt"

func main() {
	testMatrix := [][]int{
		{1, 2, 0, 3, 4, 6},
		{1, 0, 0, 4, 6, 9},
		{1, 0, 3, 0, 6, 0},
	}

	for _, v := range testMatrix {

		fmt.Print(v, " -> ")
		moveZeroes(v)
		fmt.Println(v)
	}
}

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); i, j = i+1, j+1 {
		for ; nums[j] == 0 && j < len(nums)-1; j++ {
		}
		nums[i] = nums[j]

		if i != j {
			nums[j] = 0
		}
	}
}
