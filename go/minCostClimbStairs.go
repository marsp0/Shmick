package main

func minCostClimbingStairs(cost []int) (result int) {
	// https://leetcode.com/problems/min-cost-climbing-stairs/description/
	var (
		costMap = map[int]int{}
	)
	costMap[0] = cost[0]
	costMap[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		var (
			twoSteps = costMap[i-2]
			oneStep  = costMap[i-1]
		)
		if twoSteps < oneStep {
			costMap[i] = twoSteps + cost[i]
		} else {
			costMap[i] = oneStep + cost[i]
		}
	}
	if costMap[len(cost)-1] < costMap[len(cost)-2] {
		result = costMap[len(cost)-1]
	} else {
		result = costMap[len(cost)-2]
	}
	return result
}
