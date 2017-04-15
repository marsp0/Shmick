package main

import "fmt"
import "strings"

func reverseWords(s string) string {
    var word_list = strings.Split(s, " ")
    for i := 0 ; i < len(word_list) ; i++ {
    	var temp string
    	for j := len(word_list[i])-1 ; j > -1 ; j-- {
    		temp += string(word_list[i][j])
    	}
    	word_list[i] = temp
    }
    return strings.Join(word_list, " ")
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}