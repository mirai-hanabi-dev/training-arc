package main

func findMin(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mi := lo + (hi-lo)/2

		if nums[lo] > nums[mi] {
			lo++
			hi = mi
		} else if nums[mi] > nums[hi] {
			lo = mi + 1
		} else {
			hi--
		}
	}

	return nums[lo]
}
