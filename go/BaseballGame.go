package main

import (
	"strconv"
)

func calPoints(ops []string) int {
	var (
		validRounds []int
	)
	for _, value := range ops {
		var points, err = strconv.Atoi(value)
		if err != nil {
			if value == "+" {
				var size = len(validRounds)
				validRounds = append(validRounds, validRounds[size-1]+validRounds[size-2])
			} else if value == "C" {
				validRounds = validRounds[:len(validRounds)-1]
			} else {
				validRounds = append(validRounds, validRounds[len(validRounds)-1]*2)
			}
		} else {
			validRounds = append(validRounds, points)
		}
	}
	var sum int
	for _, value := range validRounds {
		sum += value
	}
	return sum
}
