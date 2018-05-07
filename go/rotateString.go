package main

func rotateString(A string, B string) bool {
	// https://leetcode.com/problems/rotate-string/description/
	if len(A) != len(B) {
		return false
	}
	if len(A) > 0 {
		for i := 0; i < len(A); i++ {
			A = rotate(A)
			if A == B {
				return true
			}
		}
		return false
	}
	return true
}

func rotate(word string) string {
	return word[1:] + string(word[0])
}
