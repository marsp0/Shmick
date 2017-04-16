func singleNumber(nums []int) int {
    var hash = make(map[int]int)
    var result int
    for _,v := range nums {
    	hash[v] += 1
    }
    for _,v := range nums {
    	if hash[v] == 1 {
    		result = v
    	}
    }
    return result
}