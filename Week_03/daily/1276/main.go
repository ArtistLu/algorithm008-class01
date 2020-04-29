package main

import "fmt"

func main() {
	test := [][]int{
		{16, 7},
		{4, 17},
	}

	for _, args := range test {
		fmt.Println(args[0], " + ", args[1], numOfBurgers(args[0], args[1]))
	}
}

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	i, j := (tomatoSlices - 2*cheeseSlices), (4*cheeseSlices - tomatoSlices)
	if i%2 != 0 || j%2 != 0 || i < 0 || j < 0 {
		return []int{}
	}

	return []int{i / 2, j / 2}
}
