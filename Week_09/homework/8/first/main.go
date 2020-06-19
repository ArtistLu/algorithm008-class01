package main

import "fmt"

func main() {
	test := []string {
		"123",
		"-12a",
		"+-12",
		"-+12",
		"+12",
		"aa",
		" -42",
	}

	for _, t := range test {
		fmt.Println(myAtoi(t))
	}
}

func myAtoi(str string) int {
	const (
		int32Max = 1 << 32 - 1
		int32Min = -1 << 32
	)

	i, f, opt, re := 0, 0, 1, 0
	for ; f < len(str); f++ {
		if str[i] != ' ' {
			break
		}
	}

	for i = f; i < len(str); i++ {
		if i == f && str[i] == '-' {
			opt = -1
		}

		if i == f && (str[i] == '-' || str[i] == '+') {
			continue
		}

		if str[i] > '9' || str[i] < '0' {
			break
		}

		re = re * 10 + int(str[i] - '0')
		if re * opt > int32Max {
			return int32Max
		}

		if re * opt < int32Min {
			return int32Min
		}
	}

	return opt * re
}
