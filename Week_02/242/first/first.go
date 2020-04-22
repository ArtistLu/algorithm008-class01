package main

import "fmt"

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

	sb := []byte(s)
	tb := []byte(t)
	bc := make([]int, 26)
	for i, c := range sb {
		bc[c-'a']++
		bc[tb[i]-'a']--
	}

	for _, c := range bc {
		if c != 0 {
			return false
		}
	}

	return true
}
