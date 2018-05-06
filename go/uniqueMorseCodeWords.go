package main

func uniqueMorseRepresentations(words []string) (result int) {
	// https://leetcode.com/problems/unique-morse-code-words/description/
	var (
		countMap = map[string]int{}
		morseMap = map[string]string{}
		morse    = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
		alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	)
	for i, val := range alphabet {
		morseMap[val] = morse[i]
	}
	for _, word := range words {
		var morseWord string
		for _, letter := range word {
			morseWord = morseWord + morseMap[string(letter)]
		}
		countMap[morseWord]++
	}
	return len(countMap)
}
