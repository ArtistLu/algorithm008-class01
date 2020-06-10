package main

import "fmt"

func main() {
	var num uint32 = 10
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	num = num&0x55555555 + (num>>1)&0x55555555
	num = num&0x33333333 + (num>>2)&0x33333333
	num = num&0x0f0f0f0f + (num>>4)&0x0f0f0f0f
	return int(num * 0x01010101 >> 24)
}
