package main

import (
	"fmt"
	"math"
)

func main() {
	// 	[4,-1,3, 7, 9, -2, 3, -2, 4, -1,-1,-1,9]
	// [[0,1]]
	test := [][]int{
		{4, -1, 4, -2, 4, -1, -1, 3},
	}
	re := robotSim(test[0], [][]int{{2, 4}})
	fmt.Println(re)
}

func robotSim(commands []int, obstacles [][]int) int {
	directs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	sbs := make(map[int]int)
	for _, ob := range obstacles {
		sbs[ob[0]*100000+ob[1]] = 1
	}

	d, s := 0, 0
	here := []int{0, 0}
	var re float64
	for _, c := range commands {
		if c == -1 {
			d += 1
		}
		if c == -2 {
			d -= 1
		}

		if c < 0 {
			s = d % 4
			if s < 0 {
				s += 4
			}
		}

		if c > 0 {
			for i := 0; i < c; i++ {
				onestep := []int{here[0] + directs[s][0], here[1] + directs[s][1]}

				if meetObstacles(onestep, sbs) {
					break
				}
				here[0], here[1] = onestep[0], onestep[1]
			}
			re = math.Max(re, float64(here[0]*here[0]+here[1]*here[1]))
		}
	}

	return int(re)
}

func meetObstacles(step []int, obstacles map[int]int) bool {
	if _, ok := obstacles[step[0]*100000+step[1]]; ok {
		return true
	}
	return false
}
