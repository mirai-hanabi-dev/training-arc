package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	res := 0
	for i := 0; i < n; i++ {
		dp[i][0] = 0
		dp[i][1] = nums[i]
		for j := 0; j < i; j++ {
			dp[i][0] = max(dp[i][0], dp[j][0], dp[j][1])
			dp[i][1] = max(dp[i][1], dp[j][0]+nums[i])
			if i-j >= 2 {
				dp[i][1] = max(dp[i][1], dp[j][1]+nums[i])
			}
		}
		res = max(res, dp[i][0], dp[i][1])
	}

	return res
}

func max(nums ...int) int {
	res := 0
	for _, num := range nums {
		if res < num {
			res = num
		}
	}
	return res
}

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}
