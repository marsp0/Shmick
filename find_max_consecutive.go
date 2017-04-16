package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var counter int
    for _,v := range nums {
    	if v == 1 {
    		counter += 1
    	} else if v == 0 {
    		if counter > max {
    			max = counter
    		}
    		counter = 0
    	}
    }
    if counter > max {
    	max = counter
    }
    return max
}

func main() {
	var ints = []int{1,0,1,1,0,1}
	fmt.Println(findMaxConsecutiveOnes(ints))
}
