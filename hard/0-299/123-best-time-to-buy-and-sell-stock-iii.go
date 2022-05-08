package main

func maxProfit(prices []int) int {
	// dp[i][j]: profit to make if use at most i transactions at day j
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, len(prices))
	}

	for k := 1; k < 3; k++ {
		minBuy := prices[0]
		for i := 1; i < len(prices); i++ {
			minBuy = min(minBuy, prices[i]-dp[k-1][i-1])
			dp[k][i] = max(dp[k][i-1], prices[i]-minBuy)
		}
	}

	return dp[2][len(prices)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
