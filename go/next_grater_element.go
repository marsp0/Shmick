package main

func nextGreaterElement(findNums []int, nums []int) []int {
	var result []int
	var flag bool
	var flag2 bool
	for _, v := range findNums {
		flag = false
		flag2 = false
		for _, k := range nums {
			if k == v {
				flag = true
			}
			if flag {
				if k > v {
					result = append(result, k)
					flag2 = true
					break
				}
			}
		}
		if flag2 == false {
			result = append(result, -1)
		}
	}
	return result
}
