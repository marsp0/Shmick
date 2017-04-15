package main

import "fmt"

func reverseString(s string) string {
	var result string
	for i := len(s) - 1; i > -1; i-- {
		result += s[i : i+1]
	}
	return result
}

func main() {
	fmt.Println(reverseString("hello"))
}
