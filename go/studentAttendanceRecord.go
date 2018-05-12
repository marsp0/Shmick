package main

func checkRecord(s string) bool {
	// https://leetcode.com/problems/student-attendance-record-i/description/
	var (
		lateCounter   int
		absentCounter int
	)
	for _, val := range s {
		if string(val) == "A" {
			if absentCounter > 0 {
				return false
			}
			absentCounter++
			lateCounter = 0
		}
		if string(val) == "L" {
			if lateCounter == 2 {
				return false
			}
			lateCounter++
		}
		if string(val) == "P" {
			lateCounter = 0
		}
	}
	return true
}
