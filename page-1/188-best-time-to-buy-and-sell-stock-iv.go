package main

func maxProfit(k int, prices []int) int {
	// dp[i][j]: profit to make if use at most i transactions at day j
	// dp[i][j] = max(dp[i-1][k-1] + price[j] - price[k]) with k=1->n
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < k+1; i++ {
		maxBuy := -prices[0]
		for j := 1; j < n; j++ {
			maxBuy = max(maxBuy, dp[i-1][j-1]-prices[j])
			dp[i][j] = max(dp[i][j-1], maxBuy+prices[j])
		}
	}
	return dp[k][n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
