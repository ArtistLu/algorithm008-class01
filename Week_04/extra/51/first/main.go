package main

import "fmt"

func main() {
	// m := make(map[int]bool)

	// m[1] = true
	// m[9] = false
	// m[-1] = true
	// delete(m, 9)
	// fmt.Println(m)

	fmt.Println(solveNQueens(7))
}

func solveNQueens(n int) [][]string {
	result := [][]int{}
	var inner func(r []int, col, diff, sum map[int]bool)
	inner = func(r []int, col, diff, sum map[int]bool) {
		l := len(r)
		if l == n {
			result = append(result, append([]int{}, r...))
			return
		}

		for i := 0; i < n; i++ {
			if !col[i] && !diff[l-i] && !sum[l+i] {
				col[i], diff[l-i], sum[l+i] = true, true, true
				inner(append(r, i), col, diff, sum)
				delete(col, i)
				delete(diff, l-i)
				delete(sum, l+i)
			}
		}
	}

	inner([]int{}, map[int]bool{}, map[int]bool{}, map[int]bool{})
	return printQueue(result)
}

func printQueue(queue [][]int) [][]string {
	re := [][]string{}
	for _, r := range queue {
		ans := []string{}
		for _, q := range r {
			s := ""
			for i := 0; i < len(r); i++ {
				if i != q {
					s += "."
				} else {
					s += "Q"
				}
			}
			ans = append(ans, s)
		}
		re = append(re, ans)
	}

	return re
}
