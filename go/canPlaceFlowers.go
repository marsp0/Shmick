package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	// https://leetcode.com/problems/can-place-flowers/description/
	var (
		index   int
		counter int
	)
	if len(flowerbed) == 0 {
		if n > 0 {
			return false
		}
		return true
	}
	if len(flowerbed) == 1 {
		if flowerbed[0] == 1 {
			if n > 0 {
				return false
			}
			return true
		}
		return true
	}
	if n == 0 {
		return true
	}
	for index < len(flowerbed) {
		if flowerbed[index] == 0 {
			if index == 0 {
				if flowerbed[index+1] == 0 {
					counter++
					index++
				}
			} else if index == len(flowerbed)-1 {
				if flowerbed[index-1] == 0 {
					counter++
					index++
				}
			} else {
				if flowerbed[index-1] == 0 && flowerbed[index+1] == 0 {
					counter++
					index++
				}
			}

		}
		index++
	}
	return n <= counter
}
