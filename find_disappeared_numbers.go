package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
    var hash = make(map[int]int)
    var result []int
    for _,v := range nums {
    	hash[v] += 1
    }
    for i:= 1 ;i < len(nums)+1 ; i++ {
    	if hash[i] == 0 {
    		result = append(result, i)
    	}
    }
    return result
}

func main() {
	var ints = []int{1,2,1,2,3}
	fmt.Println(findDisappearedNumbers(ints))
}
