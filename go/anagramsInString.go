package main

func findAnagrams(s string, p string) []int {
	// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
	var (
		resultMap = map[int]bool{}
		letterMap = map[string]bool{}
	)
	for _, value := range p {
		letterMap[string(value)] = true
	}
	for i := 0; i < len(s)-len(p)+1; i++ {
		var _, ok = letterMap[s[i:i+1]]
		if ok {
			if i > 0 && resultMap[i-1] == true && s[i+len(p)-1:i+len(p)] == s[i-1:i] {
				resultMap[i] = true
				continue
			}
			if isAnagramAux(s[i:i+len(p)], p) {
				resultMap[i] = true
			}
		}
	}
	var (
		result = make([]int, len(resultMap))
		index  int
	)
	for key, value := range resultMap {
		if value {
			result[index] = key
		}
		index++
	}
	return result
}

func isAnagramAux(first, second string) bool {
	// had to change this code as the linear search was too slow for one of the cases
	var xor, firstSquareSum, secondSquareSum rune
	for _, char := range first {
		xor ^= char
		firstSquareSum += char * char
	}
	for _, char := range second {
		xor ^= char
		secondSquareSum += char * char
	}
	return xor == 0 && firstSquareSum == secondSquareSum
}
