package main

import (
	"strconv"
)

func selfDividingNumbers(left int, right int) (result []int) {
	// https://leetcode.com/problems/self-dividing-numbers/description/
	var i = left
	for i <= right {
		var word = strconv.Itoa(i)
		var isDividing = true
		for _, val := range word {
			var number, _ = strconv.Atoi(string(val))
			if number == 0 || i%number != 0 {
				isDividing = false
			}
		}
		if isDividing {
			result = append(result, i)
		}
		i++
	}
	return result
}
