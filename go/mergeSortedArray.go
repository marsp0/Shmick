package main

func mergeSortedArray(nums1 []int, lenFirst int, nums2 []int, lenSecond int) {
	// https://leetcode.com/problems/merge-sorted-array/description/
	var (
		firstindex  = lenFirst
		secondIndex = len(nums2) - 1
		resultIndex = len(nums1) - 1
	)

	if lenSecond > 0 {
		// fmt.Println(len(nums1), len(nums2))
		if lenFirst != 0 {
			firstindex = lenFirst - 1
		} else {
			firstindex = -1
		}
		for resultIndex >= 0 {
			// fmt.Println(resultIndex, firstindex, secondIndex, nums1)
			if secondIndex < 0 {
				break
			}
			if firstindex < 0 || nums2[secondIndex] > nums1[firstindex] {
				// fmt.Println("path 1")
				nums1[resultIndex], nums2[secondIndex] = nums2[secondIndex], nums1[resultIndex]
				secondIndex--
			} else {
				// fmt.Println("Path 2")
				nums1[resultIndex], nums1[firstindex] = nums1[firstindex], nums1[resultIndex]
				firstindex--
			}
			resultIndex--
		}
	}

}
