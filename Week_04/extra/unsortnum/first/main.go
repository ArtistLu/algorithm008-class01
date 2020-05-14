package main

import "fmt"

func main() {
	test := [][]int{
		{1, 2, 3, 4},
		{2, 1, 4, 3},
	}

	for _, t := range test {
		fmt.Println(t, "->", unSortNum(t))
	}
}

func unSortNum(num []int) int {
	re := 0
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[j] < num[i] {
				re++
			}
		}
	}

	return re * 2
}
