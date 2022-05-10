package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	i := 0
	start := intervals[0][0]
	end := intervals[0][1]

	for i < len(intervals) {
		if i == len(intervals)-1 {
			result = append(result, []int{start, end})
			break
		}

		if end < intervals[i+1][0] {
			result = append(result, []int{start, end})
			i += 1
			start = intervals[i][0]
			end = intervals[i][1]
			continue
		}

		i += 1
		end = max(end, intervals[i][1])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
