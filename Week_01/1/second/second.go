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
		fmt.Print(nums, " exists ")
		indexs := twoSum(nums, target)
		fmt.Print(nums[indexs[0]], " + ", nums[indexs[1]], " = ")
		fmt.Println(target)
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := m[target-num]; ok {
			return []int{index, i}
		}
		m[num] = i
	}

	return []int{}
}
