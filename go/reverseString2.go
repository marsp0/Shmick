package main

import (
	"strings"
)

func reverseStr(s string, k int) (result string) {
	// https://leetcode.com/problems/reverse-string-ii/description/
	var (
		counter    = 0
		stringList = strings.Split(s, "")
		index      = 0
	)
	for index < len(stringList) {
		if index == counter*2*k {
			var (
				i = counter * 2 * k
				j = i + k - 1
			)
			index = j + k
			if j >= len(stringList) {
				j = len(stringList) - 1
			}
			for i < j {
				stringList[i], stringList[j] = stringList[j], stringList[i]
				i++
				j--
			}
			counter++
		}
		index++
	}
	return strings.Join(stringList, "")
}
