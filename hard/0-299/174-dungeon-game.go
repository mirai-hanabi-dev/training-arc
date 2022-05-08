package main

import (
	"math"
)

type Position struct {
	x, y, currentHP int
}

func possibleHP(dungeon [][]int, initHP int) bool {
	m := len(dungeon)
	n := len(dungeon[0])

	currentHP := initHP + dungeon[0][0]

	if currentHP <= 0 {
		return false
	}

	vx := []int{0, 1}
	vy := []int{1, 0}

	dp := make([][]Position, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]Position, n)
	}

	start := Position{x: 0, y: 0, currentHP: currentHP}
	queue := []Position{start}
	dp[0][0].currentHP = currentHP

	for len(queue) > 0 {
		nextQueue := make([]Position, 0)
		for _, pos := range queue {
			for k := 0; k < len(vx); k++ {
				nextX := pos.x + vx[k]
				nextY := pos.y + vy[k]
				if nextX == m || nextY == n {
					continue
				}
				nextHP := dp[pos.x][pos.y].currentHP + dungeon[nextX][nextY]
				if nextHP <= 0 || nextHP <= dp[nextX][nextY].currentHP {
					continue
				}

				// If first time update, then push to queue.
				// Else skip to avoid add multiple of same cell
				// into queue.
				if dp[nextX][nextY].currentHP == 0 {
					nextQueue = append(nextQueue, Position{
						x: nextX,
						y: nextY,
					})
				}
				dp[nextX][nextY].currentHP = nextHP
			}
		}
		queue = nextQueue
	}
	return dp[m-1][n-1].currentHP > 0
}

func calculateMinimumHP(dungeon [][]int) int {
	bestHP := math.MaxInt32

	lo := 1
	hi := 200 * 200 * 1000
	for lo <= hi {
		mi := lo + (hi-lo)/2
		if possibleHP(dungeon, mi) {
			bestHP = mi
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}

	return bestHP
}
