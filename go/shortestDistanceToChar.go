package main

func shortestToChar(letters string, target byte) (result []int) {
	// https://leetcode.com/problems/shortest-distance-to-a-character/description/
	var (
		targetIndices []int
		next          int
		prev          int
		nextIndex     int
	)
	for index, val := range letters {
		if string(val) == string(target) {
			targetIndices = append(targetIndices, index)
		}
	}
	next = targetIndices[0]
	nextIndex = 0
	prev = -20000
	for index := range letters {
		if index == next {
			prev = next
			if nextIndex+1 < len(targetIndices) {
				next = targetIndices[nextIndex+1]
				nextIndex++
			} else {
				next = 20000
			}
			result = append(result, 0)
		} else {
			var (
				fromPrev = index - prev
				fromNext = next - index
			)
			if fromPrev < fromNext {
				result = append(result, fromPrev)
			} else {
				result = append(result, fromNext)
			}
		}
	}
	return result
}
