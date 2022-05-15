package main

import (
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	dis := make([]int, n)
	for idx := range dis {
		dis[idx] = math.MaxInt
	}
	dis[k-1] = 0
	for count := 0; count < n-1; count++ {
		for _, edge := range times {
			u := edge[0] - 1
			v := edge[1] - 1
			w := edge[2]
			if dis[u] < math.MaxInt && dis[u]+w < dis[v] {
				dis[v] = dis[u] + w
			}
		}
	}

	res := math.MinInt
	for _, t := range dis {
		if t > res {
			res = t
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
