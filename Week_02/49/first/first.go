package main

import ( 
	"fmt"
	"sort"
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
	var ans [][]string
	if len(strs) == 0 {
		return ans
	}

	m := make(map[string][]string, 0)
	for _, s := range strs {
		key := sortStr(s)
		m[key] = append(m[key], s)
	}
	
	for _, s := range m {
		ans = append(ans, s)
	}
	return ans
}

func sortStr(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i,j int) bool {
		return b[i] > b[j]
	})

	return string(b)
}
