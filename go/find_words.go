package main

import "strings"

func findWords(words []string) []string {
	var dict = make(map[string]int)
	var result []string
	dict["q"] = 1
	dict["w"] = 1
	dict["e"] = 1
	dict["r"] = 1
	dict["t"] = 1
	dict["y"] = 1
	dict["u"] = 1
	dict["i"] = 1
	dict["o"] = 1
	dict["p"] = 1
	dict["w"] = 1
	dict["a"] = 2
	dict["s"] = 2
	dict["d"] = 2
	dict["f"] = 2
	dict["g"] = 2
	dict["h"] = 2
	dict["j"] = 2
	dict["k"] = 2
	dict["l"] = 2
	dict["z"] = 3
	dict["x"] = 3
	dict["c"] = 3
	dict["v"] = 3
	dict["b"] = 3
	dict["n"] = 3
	dict["m"] = 3
	var flag bool
	for i := 0; i < len(words); i++ {
		var stringer = strings.ToLower(words[i][0:1])
		var currentWord = strings.ToLower(words[i])
		var currentRow = dict[stringer]
		flag = true
		for j := 1; j < len(words[i]); j++ {
			if dict[currentWord[j:j+1]] != currentRow {
				flag = false
				break
			}
		}
		if flag == true {
			result = append(result, words[i])
		}
	}
	return result
}
