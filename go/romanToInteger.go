package main

import "fmt"

func romanToInt(s string) (result int) {
	// https://leetcode.com/problems/roman-to-integer/description/
	var (
		romanLetters = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
		index        = 0
	)
	for index < len(s) {
		var stringLetter = string(s[index])
		fmt.Println(result)
		switch stringLetter {
		case "I":
			if index+1 < len(s) {
				if string(s[index+1]) == "V" {
					result += 4
					index += 2
				} else if string(s[index+1]) == "X" {
					result += 9
					index += 2
				} else {
					result++
					index++
				}
			} else {
				result += romanLetters[stringLetter]
				index++
			}
		case "X":
			if index+1 < len(s) {
				if string(s[index+1]) == "L" {
					result += 40
					index += 2
				} else if string(s[index+1]) == "C" {
					result += 90
					index += 2
				} else {
					result += romanLetters[stringLetter]
					index++
				}
			} else {
				result += romanLetters[stringLetter]
				index++
			}
		case "C":
			fmt.Println(romanLetters[stringLetter])
			if index+1 < len(s) {
				if string(s[index+1]) == "D" {
					result += 400
					index += 2
				} else if string(s[index+1]) == "M" {
					result += 900
					index += 2
				} else {
					result += romanLetters[stringLetter]
					index++
				}
			} else {
				result += romanLetters[stringLetter]
				index++
			}
		default:
			result += romanLetters[stringLetter]
			index++
		}
	}
	return result
}
