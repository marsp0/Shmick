//https://leetcode.com/problems/longest-uncommon-subsequence-i/#/description

package main

func findLUSlength(a string, b string) int {
	if len(a) != len(b) {
		if len(a) > len(b) {
			return len(a)
		}
		return len(b)
	}
	if a != b {
		return len(b)
	}
	return -1
}
