package main

import (
	"fmt"
	"math"
)

func fromIntToSlice(x int) []int {
	nums := []int{}
	for x > 0 {
		digit := x % 10
		nums = append(nums, digit)
		x = x / 10
	}
	return nums
}

func buildReverse(arr []int) int {
	num := 0
	times := 1
	for idx := range arr {
		num += arr[len(arr)-idx-1] * times
		times *= 10
	}

	return num
}

func reversePositive(x int) int {
	arrX := fromIntToSlice(x)
	arrMax := fromIntToSlice(math.MaxInt32)
	if len(arrX) < len(arrMax) {
		return buildReverse(arrX)
	}
	if len(arrX) > len(arrMax) {
		return 0
	}

	flag := true
	for idx := range arrX {
		if arrX[idx] > arrMax[len(arrX)-idx-1] {
			flag = false
			break
		} else if arrX[idx] < arrMax[len(arrX)-idx-1] {
			break
		}
	}

	if !flag {
		return 0
	}

	return buildReverse(arrX)
}

func reverse(x int) int {
	if x == math.MinInt32 {
		return 0
	}
	if x < 0 {
		return -reversePositive(-x)
	}
	return reversePositive(x)
}

func main() {
	fmt.Println(reverse(-2147483640))
}
