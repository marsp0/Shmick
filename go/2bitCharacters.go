package main

func isOneBitCharacter(bits []int) bool {
	// https://leetcode.com/problems/1-bit-and-2-bit-characters/description/
	var index = 0
	for index < len(bits) {
		if bits[index] == 1 {
			if index+1 < len(bits) && bits[index+1] == 1 || bits[index+1] == 0 {
				if index+1 == len(bits)-1 {
					return false
				}
				index += 2
			}
		} else {
			index++
		}
	}
	return true
}
