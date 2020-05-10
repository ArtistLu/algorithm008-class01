package main

import (
	"fmt"
	"sort"
)

func main() {
	test := [][][]int{
		{{5, 3, 2, 3, 1}, {3, 4, 6, 4}},
		{{2, 3, 1}, {1, 1}},
		{{2, 3, 1}, {}},
		{{10, 9, 8, 7}, {5, 6, 7, 8}},
	}
	for _, ags := range test {

		fmt.Println(ags[0], "+", ags[1], "->", findContentChildren(ags[0], ags[1]))
	}
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i int
	for j := 0; j < len(s) && i < len(g); j++ {
		if s[j] >= g[i] {
			i++
		}
	}

	return i
}
