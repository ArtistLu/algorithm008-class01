package main

import "fmt"

func main() {
	bs := "1234"
	re := reverseStr(bs, 4)
	fmt.Println(re)
}

func reverseStr(s string, k int) string {
	rs := []rune(s)
	for cur := 0; cur < len(rs); cur += 2 * k {
		end := cur + k - 1
		if end > len(rs)-1 {
			subReverse(rs, cur, len(rs)-1)
		} else {
			subReverse(rs, cur, end)
		}

	}
	return string(rs)
}

func subReverse(s []rune, l, r int) {
	for ; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
