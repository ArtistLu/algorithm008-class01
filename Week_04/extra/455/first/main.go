package main

import "sort"

func main() {
	
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
