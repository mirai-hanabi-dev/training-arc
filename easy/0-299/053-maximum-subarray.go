package main

import "math"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for idx, value := range nums {
		if idx == 0 {
			dp[idx] = value
			continue
		}
		dp[idx] = int(math.Max(float64(dp[idx-1]+nums[idx]), float64(nums[idx])))
	}

	res := math.MinInt
	for _, value := range dp {
		if value > res {
			res = value
		}
	}
	return res
}
