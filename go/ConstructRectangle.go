package main

import (
	"math"
)

func constructRectangle(area int) []int {
	// https://leetcode.com/problems/construct-the-rectangle/description/
	var current = int(math.Sqrt(float64(area)))
	for {
		if area%current == 0 {
			break
		}
		current--
	}
	var length = current
	var width = area / current
	if length < width {
		width, length = length, width
	}
	return []int{length, width}
}
