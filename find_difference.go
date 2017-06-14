//https://leetcode.com/problems/find-the-difference/#/description

package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("a","aa"))
}


func findTheDifference(s string, t string) byte {
	//assume that we only deal with ascii chars
	var hash_table = map[string]int{}
	var hash_table_t = map[string]int{}
	var result byte
	for i := 0; i < len(s); i++ {
		var _,err = hash_table[string(s[i])]
		if !err {
			hash_table[string(s[i])] = 1
			} else {
				hash_table[string(s[i])] += 1
			}
	}

	for i := 0; i < len(t); i++ {
		var _,err = hash_table_t[string(t[i])]
		if !err {
			hash_table_t[string(t[i])] = 1
			} else {
				hash_table_t[string(t[i])] += 1
			}
	}

	for i := 0 ; i < len(t) ; i++ {
		if hash_table[string(t[i])] != hash_table_t[string(t[i])] {
			result = t[i]
		}
	}
	return result
}