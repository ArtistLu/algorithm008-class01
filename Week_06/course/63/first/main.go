package main

import "fmt"

func main() {
	test := [][][]int{
		{{0}, {0}},
		{{1}, {0}},
		{
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
		},
		{
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
		},
	}

	for _, t := range test {
		fmt.Println(t, "->", uniquePathsWithObstacles(t))
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	status := make([]int, len(obstacleGrid[0]))
	for i := 0; i < len(obstacleGrid); i++ {

		if i > 0 && status[0] == 0 {

		} else {
			status[0] = 1 - obstacleGrid[i][0]
		}

		for j := 1; j < len(status); j++ {
			if obstacleGrid[i][j] == 1 {
				status[j] = 0
				continue
			}
			status[j] += status[j-1]
		}
	}

	return status[len(status)-1]
}
