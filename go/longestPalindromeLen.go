package main

func longestPalindrome(s string) int {
	// https://leetcode.com/problems/longest-palindrome/description/
	var (
		letterFrequency = map[string]int{}
		result          int
		singleChar      = true
	)
	for _, value := range s {
		letterFrequency[string(value)]++
	}
	for _, value := range letterFrequency {
		if value > 1 {
			if value%2 == 0 {
				result += value
			} else {
				if singleChar {
					result += value
					singleChar = false
				} else {
					result += (value - 1)
				}
			}
		} else {

			if singleChar {
				result++
				singleChar = false
			}
		}
	}
	return result
}
