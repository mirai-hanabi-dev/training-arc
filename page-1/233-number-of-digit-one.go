package main

import (
	"math"
)

func dfsCountDigitOne(n int, cache map[int]int) int {
	if n == 0 {
		return 0
	}
	if v, ok := cache[n]; ok {
		return v
	}

	tmp := n
	var firstDigit int
	var restDigit int
	var length int

	for tmp > 0 {
		if tmp < 10 {
			firstDigit = tmp
		}
		length += 1
		tmp /= 10
	}
	restDigit = n - firstDigit*int(math.Pow(10, float64(length-1)))

	count := 0
	mask := int(math.Pow(10, float64(length-1))) - 1
	for i := 0; i < firstDigit; i++ {
		count += dfsCountDigitOne(mask, cache)
		if i == 1 {
			count += (mask + 1)
		}
	}

	count += dfsCountDigitOne(restDigit, cache)
	if firstDigit == 1 {
		count += (restDigit + 1)
	}

	cache[n] = count
	return count
}

func countDigitOne(n int) int {
	cache := map[int]int{}
	return dfsCountDigitOne(n, cache)
}
