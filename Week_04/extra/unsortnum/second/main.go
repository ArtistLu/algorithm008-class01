package main

import "fmt"

func main() {
	test := [][]int{
		{1, 3, 8, 4, 7, 3, 8},
		{1, 19, 8, 4, 7, 3, 8},
		{1, 3, 8, 9, 7, 3, 8},
		{1, 3, 18, 4, 7, 3, 8},
	}
	for _, t := range test {

		fmt.Println(t, "unsort num ", mergeSort(append([]int{}, t...)))
	}
	fmt.Println(test)
}

func mergeSort(nums []int) int {
	var unsort int
	innerSort(nums, 0, len(nums)-1, &unsort)
	return unsort
}

func innerSort(nums []int, l, r int, unsort *int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1

	innerSort(nums, l, mid, unsort)
	innerSort(nums, mid+1, r, unsort)

	a, b, c := nums[l:mid+1], nums[mid+1:r+1], make([]int, r-l+1)
	var i, j, k int
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c[k] = a[i]
			i, k = i+1, k+1
		} else {
			*unsort += len(a) - i
			c[k] = b[j]
			j, k = j+1, k+1
		}
	}

	for i < len(a) {
		c[k] = a[i]
		i, k = i+1, k+1
	}

	for j < len(b) {
		c[k] = b[j]
		j, k = j+1, k+1
	}

	k = 0
	for i := l; i <= r; i++ {
		nums[i], k = c[k], k+1
	}
}
