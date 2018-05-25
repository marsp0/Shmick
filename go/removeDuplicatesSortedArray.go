package main

func removeDuplicates(nums []int) int {
	var (
		lastAvailableIndex = len(nums) - 1
		i                  = 1
		j                  int
		current            int
	)
	if len(nums) >= 1 {
		current = nums[0]
	} else {
		return 0
	}
	for i < len(nums) && lastAvailableIndex > i {
		if nums[i] != current {
			current = nums[i]
		} else {
			nums[i], nums[lastAvailableIndex] = nums[lastAvailableIndex], nums[i]
			j = lastAvailableIndex - 1
			for nums[j] == nums[i] && j >= i {
				// fmt.Println("ds")
				j--

			}
			lastAvailableIndex = j
		}
		i++
	}
	return lastAvailableIndex + 1
}
