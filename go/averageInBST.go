package main

func averageOfLevels(root *TreeNode) []float64 {
	// https://leetcode.com/problems/average-of-levels-in-binary-tree/
	var (
		result     = []int{}
		toReturn   = []float64{}
		levelCount = map[int]int{}
	)
	getAverageAux(root, &result, 0, levelCount)
	for index := range result {
		if index > 0 {
			toReturn = append(toReturn, float64(result[index])/float64(levelCount[index]))
		} else {
			toReturn = append(toReturn, float64(result[index]))
		}
	}
	return toReturn
}

func getAverageAux(root *TreeNode, result *[]int, level int, levelCount map[int]int) {
	if root != nil {
		if level >= len(*result) {
			*result = append((*result), 0)
		}
		(*result)[level] += root.Val
		levelCount[level]++
		getAverageAux(root.Left, result, level+1, levelCount)
		getAverageAux(root.Right, result, level+1, levelCount)
	}
}
