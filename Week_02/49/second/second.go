package main

import (
	"fmt"
)

func main() {
	testMatrix := [][]string{
		{"abx", "xab", "bxa", "zhang", "hangz", "lu", "ll"},
		{},
		{"zy", "h"},
	}

	for _, strs := range testMatrix {
		fmt.Println(strs, "-> ", groupAnagrams(strs))
	}

}

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	m := make(map[[26]byte][]string, 0)
	for _, s := range strs {
		h := shash(s)
		m[h] = append(m[h], s)
	}

	for _, ss := range m {
		ans = append(ans, ss)
	}

	return ans
}

func shash(s string) [26]byte {
	h := [26]byte{}

	for _, c := range s {
		h[c-'a']++
	}

	return h
}
