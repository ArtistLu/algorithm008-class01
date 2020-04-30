package main

import "fmt"

func main() {
	test := []string{
		"abc",
		"a b c",
		" ",
		"    ",
	}
	for _, s := range test {
		fmt.Println(s, "-> ", replaceSpace(s))
	}
}

func replaceSpace(s string) string {
	re := []byte{}
	b := []byte(s)
	for _, v := range b {
		if v != 32 {
			re = append(re, v)
		} else {
			re = append(re, []byte{37, 50, 48}...)
		}
	}
	return string(re)
}
