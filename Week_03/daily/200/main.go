package main

import "fmt"

func main() {
	test := [][]byte{
		{'1', '0', '1', '0'},
		{'1', '0', '1', '0'},
		{'1', '0', '1', '0'},
		{'0', '1', '1', '0'},
		{'1', '0', '1', '0'},
	}

	islands := numIslands(test)
	fmt.Println(islands)
}

const (
	land  = '1'
	water = '0'
)

func numIslands(grid [][]byte) int {
	var islands int
	for i, row := range grid {
		for j, area := range row {
			if area == water {
				continue
			}
			islands++
			fillWithWater(grid, i, j)
		}
	}

	return islands
}

func fillWithWater(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == water {
		return
	}
	grid[i][j] = water
	fillWithWater(grid, i+1, j)
	fillWithWater(grid, i-1, j)
	fillWithWater(grid, i, j+1)
	fillWithWater(grid, i, j-1)
}
