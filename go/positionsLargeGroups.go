package main

func largeGroupPositions(S string) (result [][]int) {
	// https://leetcode.com/problems/positions-of-large-groups/description/
	var (
		start int
		end   int
	)
	for index, value := range S {
		if index > 0 {
			if string(value) != string(S[index-1]) {
				if end-start >= 2 {
					result = append(result, []int{start, end})
				}
				start = index
				end = index
			} else {
				end = index
			}
		}
	}
	if end-start >= 2 {
		result = append(result, []int{start, end})
	}
	return result
}
