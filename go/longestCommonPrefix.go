package main

func longestCommonPrefix(strs []string) string {
	// https://leetcode.com/problems/longest-common-prefix/description/
	var (
		index     = 0
		counter   = 0
		current   byte
		wordIndex = 0
	)
	if len(strs) == 0 || (len(strs) >= 1 && len(strs[0]) == 0) {
		return ""
	}
	current = strs[0][0]
	for wordIndex < len(strs) {
		if index < len(strs[wordIndex]) {
			if strs[wordIndex][index] != current {
				break
			}
		} else {
			break
		}
		if wordIndex == len(strs)-1 {
			counter++
			if index+1 < len(strs[wordIndex]) {
				current = strs[wordIndex][index+1]
			} else {
				break
			}

			index++
			wordIndex = 0

			continue
		}
		wordIndex++
	}
	return strs[0][:counter]

}
