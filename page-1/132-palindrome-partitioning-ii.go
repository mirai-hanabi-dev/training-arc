package main

type Pair struct {
	i, j int
}

func minCut(s string) int {
	isPalindrome := map[Pair]bool{}

	for i := 0; i < len(s); i++ {
		left := i
		right := i
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				break
			}
			isPalindrome[Pair{left, right + 1}] = true
			left--
			right++
		}
		if i == len(s)-1 {
			continue
		}
		left = i
		right = i + 1
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				break
			}
			isPalindrome[Pair{left, right + 1}] = true
			left--
			right++
		}
	}

	dp := make([]int, len(s))

	for j := 0; j < len(s); j++ {
		if isPalindrome[Pair{0, j + 1}] {
			dp[j] = 0
			continue
		}
		dp[j] = dp[j-1] + 1
		for i := j - 1; i > 0; i-- {
			if isPalindrome[Pair{i, j + 1}] {
				dp[j] = min(dp[j], dp[i-1]+1)
			}
		}
	}

	return dp[len(s)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
