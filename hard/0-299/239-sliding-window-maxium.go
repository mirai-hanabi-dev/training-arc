package main

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, n-k+1)

	q := make([]int, n)
	l := 0
	r := -1

	for i := 0; i < k; i++ {
		for r-l >= 0 && nums[i] >= nums[q[r]] {
			r -= 1
		}
		r += 1
		q[r] = i
	}

	for i := 0; i < n-k+1; i++ {
		if q[l] == i-1 {
			l += 1
		}

		result[i] = nums[q[l]]
		if i == n-k {
			break
		}

		for r-l >= 0 && nums[i+k] >= nums[q[r]] {
			r -= 1
		}
		r += 1
		q[r] = i + k
	}

	return result
}
