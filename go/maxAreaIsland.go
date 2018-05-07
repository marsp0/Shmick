package main

func maxAreaOfIsland(grid [][]int) int {
	var (
		visited   = map[int]map[int]int{}
		perimeter = 0
	)
	for i := range grid {
		visited[i] = map[int]int{}
	}
	for i := range grid {
		for j := range grid[i] {
			var _, isVisited = visited[i][j]
			if grid[i][j] == 1 && !isVisited {
				var newPerimeter = visit(grid, i, j, visited)
				if newPerimeter > perimeter {
					perimeter = newPerimeter
				}
			}
		}
	}
	return perimeter
}

func visit(grid [][]int, i, j int, visited map[int]map[int]int) int {
	var cost = 0
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == 1 {
		var _, isVisited = visited[i][j]
		if !isVisited {
			visited[i][j] = 1
			var (
				adjacent = [][]int{{i, j - 1}, {i - 1, j}, {i, j + 1}, {i + 1, j}}
			)
			cost = 1
			for _, value := range adjacent {
				cost += visit(grid, value[0], value[1], visited)
			}
		}
	}
	return cost
}
