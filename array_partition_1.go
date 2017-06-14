//https://leetcode.com/problems/array-partition-i/#/description - easy

package main 

import "sort"
import "fmt"

func main() {
	var array  = []int{1,4,3,2}
	fmt.Println(arrayPairSum(array))
}

func arrayPairSum(nums []int) int {
	sort.Slice(nums, func (i,j int) bool {return nums[i] < nums[j]} )
	var result int
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}