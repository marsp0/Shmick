package main

func validPalindrome(s string) bool {
	// https://leetcode.com/problems/valid-palindrome-ii/description/
	var (
		i int
		j = len(s) - 1
	)
	for i < j {
		if s[i] != s[j] {
			return validPalindromeAux(s[i+1:j+1]) || validPalindromeAux(s[i:j])
		}
		i++
		j--
	}
	return true
}

func validPalindromeAux(word string) bool {
	if len(word) == 1 {
		return true
	}
	var (
		i int
		j = len(word) - 1
	)
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}
