package main

func generate(numRows int) (result [][]int) {
	// https://leetcode.com/problems/pascals-triangle/description/
	if numRows == 0 {
		return result
	} else if numRows == 1 {
		result = append(result, []int{1})
		return result
	}
	result = append(result, []int{1})
	result = append(result, []int{1, 1})
	for i := 2; i < numRows; i++ {
		result = append(result, []int{})
		for j := 0; j < len(result[i-1])+1; j++ {
			if j == 0 || j == len(result[i-1]) {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i], result[i-1][j]+result[i-1][j-1])
			}

		}
	}
	return result
}
