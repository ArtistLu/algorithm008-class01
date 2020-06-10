package main

import "fmt"

func main() {
	test := []uint32{1, 2, 3, 4, 5}
	for _, t := range test {
		fmt.Println(t, "->", reverseBits(t))
	}
}

func reverseBits(num uint32) uint32 {
	var re uint32
	for i := 1; i <= 32; i++ {
		re <<= 1
		re += num & 1
		num >>= 1
	}
	return re
}
