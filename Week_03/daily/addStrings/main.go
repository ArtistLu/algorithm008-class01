package main

import "fmt"

func main() {
	test := [][]string{
		{"123", "997"},
		{"123", "9090990"},
		{"234", "678"},
		{"", "678"},
		{"", ""},
	}

	for _, str := range test {
		fmt.Println(str[0], " + ", str[1], " = ", addStrings(str[0], str[1]))
	}
}

func addStrings(num1 string, num2 string) string {
	str1, str2 := []byte(num1), []byte(num2)
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	var sum byte
	for i, j := len(str1)-1, len(str2)-1; i >= 0; i, j = i-1, j-1 {
		if j >= 0 {
			sum += str2[j] - '0'
		}
		sum += str1[i] - '0'

		str1[i] = sum%10 + '0'
		sum /= 10
	}

	if sum > 0 {
		str1 = append([]byte{'1'}, str1...)
	}

	return string(str1)
}
