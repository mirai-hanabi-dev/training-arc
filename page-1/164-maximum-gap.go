package main

import "math"

type Bucket struct {
	min, max int
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min := math.MaxInt32
	max := -math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	if min == max {
		return 0
	}

	n := len(nums)
	size := int(math.Ceil(float64(max-min) / float64(n-1)))

	bucket := make([]Bucket, n)
	for i := 0; i < n; i++ {
		bucket[i].min = math.MaxInt32
		bucket[i].max = -math.MaxInt32
	}

	for _, num := range nums {
		i := (num - min) / size
		if num < bucket[i].min {
			bucket[i].min = num
		}
		if num > bucket[i].max {
			bucket[i].max = num
		}
	}

	res := -math.MaxInt32
	i := 0
	for {
		if i >= n {
			break
		}
		for i < n && bucket[i].max == -math.MaxInt32 {
			i++
		}
		j := i + 1
		if j >= n {
			break
		}
		for j < n && bucket[j].min == math.MaxInt32 {
			j++
		}
		if j >= n {
			break
		}
		if bucket[j].min-bucket[i].max > res {
			res = bucket[j].min - bucket[i].max
		}
		i = j
	}

	return res
}
