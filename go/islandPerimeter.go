package main

func islandPerimeter(grid [][]int) int {
	// https://leetcode.com/problems/island-perimeter/description/
	var (
		counter = 0
	)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				counter += adjacent(i, j, grid)
			}
		}
	}
	return counter
}

func adjacent(i, j int, grid [][]int) int {
	var adjacent = [][]int{{i, j - 1}, {i - 1, j}, {i, j + 1}, {i + 1, j}}
	var result = 0
	for _, value := range adjacent {
		if value[0] < 0 || value[1] < 0 || value[0] == len(grid) || value[1] == len(grid[0]) || grid[value[0]][value[1]] == 0 {
			result++
		}
	}
	return result
}
