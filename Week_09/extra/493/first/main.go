package main

import "fmt"

func main() {
	nums := []int{2, 9, 1, 3, 2, 9, 2}
	pairs := reversePairs((nums))
	fmt.Println(pairs)
}

func reversePairs(nums []int) int {
	re := 0
	mergesort(nums, 0, len(nums)-1, &re)
	return re
}

func mergesort(nums []int, s, e int, re *int) {
	if s >= e {
		return
	}

	mid := s + (e-s)>>1
	mergesort(nums, s, mid, re)
	mergesort(nums, mid+1, e, re)
	merge(nums, s, mid, e, re)
}

func merge(nums []int, s, mid, e int, re *int) {
	tmp := make([]int, e-s+1)
	t, i, j, k := s, s, mid+1, 0
	for ; j <= e; j++ {
		for ; t <= mid && nums[t] <= nums[j]<<1; t++ {
		}
		for ; i <= mid && nums[i] <= nums[j]; i, k = i+1, k+1 {
			tmp[k] = nums[i]
		}
		tmp[k], k = nums[j], k+1
		*re += mid - t + 1
	}

	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = nums[i]
	}

	for k, i := 0, s; i <= e; i, k = i+1, k+1 {
		nums[i] = tmp[k]
	}
}
