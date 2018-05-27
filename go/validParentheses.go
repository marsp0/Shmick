package main

func isValid(s string) bool {
	// https://leetcode.com/problems/valid-parentheses/description/
	var (
		stack   = []string{}
		opening = map[string]int{"(": 1, "[": 2, "{": 3}
		closing = map[string]int{")": 1, "]": 2, "}": 3}
	)
	for _, value := range s {
		var _, ok = opening[string(value)]
		if !ok {
			if len(stack) > 0 {
				if closing[string(value)] != opening[stack[len(stack)-1]] {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(value))
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
