//https://leetcode.com/problems/add-digits/#/description

package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	var result int
    var str  = strconv.Itoa(num)
    for len(str) > 1{
        var temp int
        for i := 0; i < len(str); i++ {
            s, _ := strconv.Atoi(string(str[i]))
            temp += s
        }
        str = strconv.Itoa(temp)

    }
    result, _ = strconv.Atoi(str)
    return result
}