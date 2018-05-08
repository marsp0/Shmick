package main

func canConstruct(ransomNote string, magazine string) bool {
	var (
		letterTable = map[string]int{}
	)
	for _, value := range magazine {
		letterTable[string(value)]++
	}
	for _, value := range ransomNote {
		letterTable[string(value)]--
		if letterTable[string(value)] < 0 {
			return false
		}
	}
	return true
}
