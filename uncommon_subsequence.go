//https://leetcode.com/problems/longest-uncommon-subsequence-i/#/description

package main

import "fmt"

func main() {
	fmt.Println(findLUSlength("aba","cdc"))
}


func findLUSlength(a string, b string) int {
	if len(a) != len(b) {
		if len(a) > len(b) {
			return len(a)
		} else {
			return len(b)
		}
	} else {
		if a != b {
			return len(b)
		}
		return -1
	}

}