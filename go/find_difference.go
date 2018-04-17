//https://leetcode.com/problems/find-the-difference/#/description

package main

func findTheDifference(s string, t string) byte {
	//assume that we only deal with ascii chars
	var hashTable = map[string]int{}
	var HashTableT = map[string]int{}
	var result byte
	for i := 0; i < len(s); i++ {
		var _, err = hashTable[string(s[i])]
		if !err {
			hashTable[string(s[i])] = 1
		} else {
			hashTable[string(s[i])]++
		}
	}

	for i := 0; i < len(t); i++ {
		var _, err = HashTableT[string(t[i])]
		if !err {
			HashTableT[string(t[i])] = 1
		} else {
			HashTableT[string(t[i])]++
		}
	}

	for i := 0; i < len(t); i++ {
		if hashTable[string(t[i])] != HashTableT[string(t[i])] {
			result = t[i]
		}
	}
	return result
}
