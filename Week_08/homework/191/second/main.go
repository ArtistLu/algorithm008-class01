package main

import "fmt"

func main() {
	var num uint32 = 10
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	re := 0
	for num > 0 {
		num &= (num - 1)
		re++
	}
	return re
}
