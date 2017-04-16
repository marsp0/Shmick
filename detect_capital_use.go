package main

import "fmt"
import "strings"

func detectCapitalUse(word string) bool {
	var result bool
    if strings.ToLower(word) == word {
    	result = true
    } else if strings.ToUpper(word) == word {
    	result = true
    } else {
    	if word[0:1] == strings.ToUpper(word[0:1]) {
    		if word[1:] == strings.ToLower(word[1:]) {
    			result = true
    		} else {
    			result = false
    		}
		} else {
			result = false
		}
    }
    return result
}

func main() {
	fmt.Println(detectCapitalUse("dSa"))
}
