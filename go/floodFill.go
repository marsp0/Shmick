package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// https://leetcode.com/problems/flood-fill/description/
	var (
		visited = map[int]map[int]int{}
	)
	for index := range image {
		visited[index] = map[int]int{}
	}
	colorImage(&image, sr, sc, image[sr][sc], newColor, visited)
	return image
}

func colorImage(image *[][]int, i, j, oldColor, newColor int, visited map[int]map[int]int) {
	if i >= 0 && i < len(*image) && j >= 0 && j < len((*image)[i]) && (*image)[i][j] == oldColor {
		var _, inMap = visited[i][j]
		if !inMap {
			(*image)[i][j] = newColor
			visited[i][j] = 1
			var friends = [][]int{{i, j - 1}, {i - 1, j}, {i, j + 1}, {i + 1, j}}
			for _, value := range friends {
				colorImage(image, value[0], value[1], oldColor, newColor, visited)
			}
		}
	}
}
