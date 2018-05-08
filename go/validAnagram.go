package main

func isAnagram(s string, t string) bool {
	// https://leetcode.com/problems/valid-anagram/description/
	var (
		letterTable = map[string]int{}
	)
	if len(s) != len(t) {
		return false
	}
	for _, value := range t {
		letterTable[string(value)]++
	}
	for _, value := range s {
		letterTable[string(value)]--
		if letterTable[string(value)] == 0 {
			delete(letterTable, string(value))

		}
	}
	if len(letterTable) == 0 {
		return true
	}
	return false
}
