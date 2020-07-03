package main

import (
	"alg/nextweek/kmp"
	"fmt"
)

func main() {
	// p := "abc"
	// s := "aaaaaaavabbbbcaaabcaaabc=abc="

	p := "abc"
	s := "aaaaaabc"
	kmp := kmp.NewKmp(p)
	index := kmp.Search(s)
	fmt.Println(index)

	if index == -1 {
		fmt.Println("no match")
	} else {
		fmt.Println("math ", s[index:index+len(p)])
	}
}
