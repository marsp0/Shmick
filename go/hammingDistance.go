package main

import (
	"strconv"
)

func hammingDistance(x int, y int) int {
	// https://leetcode.com/problems/hamming-distance/description/
	var (
		firstBinary  = strconv.FormatInt(int64(x), 2)
		secondBinary = strconv.FormatInt(int64(y), 2)
		counter      = 0
	)
	if len(firstBinary) < len(secondBinary) {

		var someNumber = len(secondBinary) - len(firstBinary)
		for i := 0; i < someNumber; i++ {
			firstBinary = "0" + firstBinary
		}
	} else {
		var someNumber = len(firstBinary) - len(secondBinary)

		for i := 0; i < someNumber; i++ {
			secondBinary = "0" + secondBinary
		}
	}
	for index, val := range firstBinary {
		if string(val) != string(secondBinary[index]) {
			counter++
		}
	}
	return counter
}
