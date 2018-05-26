package main

func countSegments(s string) int {
	// https://leetcode.com/problems/number-of-segments-in-a-string/description/
	var (
		result        int
		letterCounter int
		i             int
	)
	for i < len(s) {
		if string(s[i]) == " " {
			if letterCounter > 0 {
				result++
				letterCounter = 0
			}
		} else {
			letterCounter++
		}
		i++
	}
	if letterCounter > 0 {
		result++
	}
	return result
}
