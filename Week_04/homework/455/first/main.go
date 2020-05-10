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
	}
	for _, ags := range test {

		fmt.Println(ags[0], "+", ags[1], "->", findContentChildren(ags[0], ags[1]))
	}
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var re int
	var i int
	for _, c := range g {
		for ; i < len(s); i++ {
			if s[i] >= c {
				re++
				i++
				break
			}
		}

		if i > len(s)-1 {
			break
		}
	}
	return re
}
