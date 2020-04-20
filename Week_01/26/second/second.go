package main

import "fmt"

func main() {
	testMatrix := [][]int{{1, 2, 2, 3, 6}, {}, {1, 1, 1, 1}, {1, 1, 1, 1, 1, 2, 3, 4, 4, 4, 4, 5}}
	for _, v := range testMatrix {
		fmt.Print(v, "->")
		fmt.Print(removeDuplicates(v), "|", v)
		fmt.Println()
	}
}

func removeDuplicates(nums []int) int {
	var j int
	for i := 0; i < len(nums); i, j = i+1, j+1 {
		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		}
		nums[j] = nums[i]
	}
	return j
}
