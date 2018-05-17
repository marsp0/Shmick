package main

func getRow(rowIndex int) []int {
	// https://leetcode.com/problems/pascals-triangle-ii/description/
	var (
		i       = 2
		zeroRow = []int{1}
		prevRow = []int{1, 1}
	)
	for i < rowIndex+1 {
		var (
			j          int
			currentRow = []int{}
		)
		for j < len(prevRow)+1 {
			if j == 0 || j == len(prevRow) {
				currentRow = append(currentRow, 1)
			} else {
				currentRow = append(currentRow, prevRow[j]+prevRow[j-1])
			}
			j++
		}
		prevRow = currentRow
		i++
	}
	if rowIndex == 0 {
		return zeroRow
	}
	return prevRow
}
