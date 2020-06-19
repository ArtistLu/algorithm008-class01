package main

import "fmt"

func main() {
	test := []string{
		"leetcode",
		"ok",
	}

	for _, s := range test {
		fmt.Println(firstUniqChar(s))
	}
}

func firstUniqChar(s string) int {
	m := [26]int{}
	for _, c := range s {
		m[c-'a']++
	}

	for i, c := range s {
		if m[c-'a'] == 1 {
			return i
		}
	}

	return -1
}
