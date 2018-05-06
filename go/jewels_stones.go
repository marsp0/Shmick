package main

func numJewelsInStones(J string, S string) (result int) {
	// https://leetcode.com/problems/jewels-and-stones/description/
	var jewels = map[string]bool{}
	for _, val := range J {
		jewels[string(val)] = true
	}
	for _, val := range S {
		var _, ok = jewels[string(val)]
		if ok {
			result++
		}
	}
	return result
}
