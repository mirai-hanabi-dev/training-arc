package main

func numDistinct(s string, t string) int {
	// dp[i][j]: number of distinct subsequences of s[:i] which equals t[:j].
	// Result is at dp[len(s)][len(t)]

	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}

	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(t)]
}
