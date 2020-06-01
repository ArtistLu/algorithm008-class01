package main

import "fmt"

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'C', 'C', 'D', 'S', 'Z', 'E'}

	fmt.Println(leastInterval(tasks, 3))
}

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	for _, c := range tasks {
		m[c-'A']++
	}

	max, cnt := 0, 0
	for _, v := range m {
		if v == max {
			cnt++
		} else if v > max {
			max = v
			cnt = 1
		}
	}

	total := (max-1)*(n+1) + cnt

	if total > len(tasks) {
		return total
	}

	return len(tasks)
}
