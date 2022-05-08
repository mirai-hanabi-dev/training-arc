package main

import "fmt"

func reverse(nums []int, left int, right int) {
	if left >= right {
		return
	}
	n := (right - left + 1)
	for i := 0; i < n/2; i++ {
		tmp := nums[left+i]
		nums[left+i] = nums[right-i]
		nums[right-i] = tmp
	}
}

func nextPermutation(nums []int) {
	n := len(nums)
	start := n - 2
	for start >= 0 && nums[start] >= nums[start+1] {
		start--
	}

	if start == -1 {
		reverse(nums, 0, n-1)
		return
	}

	pos := start + 1
	for i := start + 1; i < n; i++ {
		if nums[i] > nums[start] && nums[i] <= nums[pos] {
			pos = i
		}
	}
	tmp := nums[start]
	nums[start] = nums[pos]
	nums[pos] = tmp

	reverse(nums, start+1, n-1)
}

func main() {
	nums := []int{2, 3, 1, 3, 3}
	for i := 0; i < 10; i++ {
		nextPermutation(nums)
		fmt.Println(nums)
	}
}
