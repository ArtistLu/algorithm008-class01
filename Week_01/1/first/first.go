package main

import "fmt"

func main() {
	testMatrix := [][]int{
		{1, 2, 3, 4, 6},
		{1, 2, 3, 4, 6, 9},
		{1, 2, 3, 4, 6, 5},
	}

	for _, v := range testMatrix {
		nums := v[0 : len(v)-1]
		target := v[len(v)-1]
		fmt.Print(nums, "->")
		fmt.Print(twoSum(nums, target), "to target ")
		fmt.Println(target)
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index := target - nums[i]
		if b, ok := m[index]; ok {
			return []int{b, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
