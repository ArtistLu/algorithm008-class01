package main

import (
	"fmt"
)

func main() {
	testMatrix := [][]string{
		{"abc", "abc"},
		{"cda", "dca"},
		{"abab", "baba"},
		{"absb", "baba"},
		{"abab", "baba"},
		{"aacc", "ccac"},
	}

	for _, v := range testMatrix {
		s := v[0]
		t := v[1]
		fmt.Println(s, "| ", t, isAnagram(s, t))
	}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	char := [26]int{}
	for _, c := range s {
		char[c-'a']++
	}

	for _, c := range t {
		char[c-'a']--
		if char[c-'a'] < 0 {
			return false
		}
	}

	return true
}
