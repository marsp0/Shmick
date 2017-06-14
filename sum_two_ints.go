// https://leetcode.com/problems/sum-of-two-integers/#/description

package main

import "fmt"

func main() {
	fmt.Println(getSum(1,2))
}

func getSum(a int, b int) int {
    if b == 0 {
    	return a
    } else { 
    	return getSum(a ^ b, (a&b) << 1 )
    }
}