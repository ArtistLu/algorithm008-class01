package main

import "fmt"

func main() {
	testMatrix := [][][]int{
		{{1}, {2}},
		{{1, 2, 3}, {2, 3, 5}},
		{{}, {2}},
		{{1, 5, 7}, {2}},
	}

	for _, v := range testMatrix {
		m := len(v[0])
		n := len(v[1])
		i := len(v[1])
		for i > 0 {
			i--
			v[0] = append(v[0], 0)
		}

		fmt.Print(v[0], "merge", v[1])
		merge(v[0], m, v[1], n)
		fmt.Print("=", v[0])
		fmt.Println()
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	i, j := m-1, n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	if j >= 0 {
		for k >= 0 {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
}
