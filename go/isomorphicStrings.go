package main

func isIsomorphic(s string, t string) bool {
	// https://leetcode.com/problems/isomorphic-strings/description/
	var (
		sTable = map[string][]int{}
		tTable = map[string][]int{}
	)
	for index, value := range s {
		var _, ok = sTable[string(value)]
		if !ok {
			sTable[string(value)] = []int{}
		}
		sTable[string(value)] = append(sTable[string(value)], index)

		_, ok = tTable[string(t[index])]
		if !ok {
			tTable[string(t[index])] = []int{}
		}
		tTable[string(t[index])] = append(tTable[string(t[index])], index)
	}
	for index, value := range s {
		for i := range sTable[string(value)] {
			if len(sTable[string(value)]) == len(tTable[string(t[index])]) {
				if sTable[string(value)][i] != tTable[string(t[index])][i] {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
