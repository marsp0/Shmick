package main

func firstUniqChar(s string) int {
	// https://leetcode.com/problems/first-unique-character-in-a-string/description/
	var (
		hashTable = map[string]int{}
	)
	for _, value := range s {
		hashTable[string(value)]++
	}
	for index, value := range s {
		if hashTable[string(value)] == 1 {
			return index
		}
	}
	return -1
}
