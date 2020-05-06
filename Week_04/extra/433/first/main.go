package main

import "fmt"

func main() {
	// 	"AAAACCCC"
	// "CCCCCCCC"
	// ["AAAACCCA","AAACCCCA","AACCCCCA","AACCCCCC","ACCCCCCC","CCCCCCCC","AAACCCCC","AACCCCCC"]
	start := "AAAACCCC"
	end := "CCCCCCCC"
	bank := []string{
		"AAAACCCA", "AAACCCCA", "AACCCCCA", "AACCCCCC", "ACCCCCCC", "CCCCCCCC", "AAACCCCC", "AACCCCCC",
	}

	fmt.Println(minMutation(start, end, bank))
}

func minMutation(start string, end string, bank []string) int {
	level, _ := bfs(start, end, bank, 0)
	return level
}

func bfs(start, end string, bank []string, level int) (int, bool) {
	if start == end {
		return level, true
	}

	if len(bank) == 0 {
		return -1, false
	}

	left := []string{}
	starts := []string{}
	for i := 0; i < len(bank); i++ {
		if isnext(start, bank[i]) {
			starts = append(starts, bank[i])
		} else {
			left = append(left, bank[i])
		}
	}

	min := -1
	for i := 0; i < len(starts); i++ {
		if v, ok := bfs(starts[i], end, left, level+1); ok {
			if min == -1 || v < min {
				min = v
			}
		}
	}

	return min, min != -1
}

func isnext(s, n string) bool {
	diff := 0
	for i := 0; i < len(s); i++ {
		if s[i] != n[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}
