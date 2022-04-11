package main

import (
	"fmt"
	"math"
)

type Hero struct {
	initHP    int
	currentHP int
	x, y      int
}

func calculateMinimumHP(dungeon [][]int) int {
	queue := make([]Hero, 0)
	m := len(dungeon)
	n := len(dungeon[0])

	bestHP := make([][]Hero, m)
	for i := 0; i < m; i++ {
		bestHP[i] = make([]Hero, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bestHP[i][j].initHP = math.MaxInt32
			bestHP[i][j].currentHP = -math.MaxInt32
		}
	}

	start := dungeon[0][0]
	var initHP, currentHP int
	if start < 0 {
		initHP = -start + 1
		currentHP = 1
	} else {
		initHP = 1
		currentHP = start + 1
	}
	queue = append(queue, Hero{initHP: initHP, currentHP: currentHP, x: 0, y: 0})
	bestHP[0][0].initHP = initHP
	bestHP[0][0].currentHP = currentHP

	vx := []int{0, 1}
	vy := []int{1, 0}

	for len(queue) > 0 {
		fmt.Println(queue)
		nextQueue := make([]Hero, 0)
		for _, hero := range queue {
			for k := 0; k < len(vx); k++ {
				nextX := hero.x + vx[k]
				nextY := hero.y + vy[k]
				if nextX == m || nextY == n {
					continue
				}

				var initHP, currentHP int
				if dungeon[nextX][nextY] >= 0 || hero.currentHP+dungeon[nextX][nextY] > 0 {
					initHP = hero.initHP
					currentHP = hero.currentHP + dungeon[nextX][nextY]
				} else {
					initHP = hero.initHP - (hero.currentHP + dungeon[nextX][nextY]) + 1
					currentHP = 1
				}
				cont1 := initHP < bestHP[nextX][nextY].initHP
				cont2 := initHP == bestHP[nextX][nextY].initHP && currentHP > bestHP[nextX][nextY].currentHP
				if cont1 || cont2 {
					nextQueue = append(nextQueue, Hero{
						initHP:    initHP,
						currentHP: currentHP,
						x:         nextX,
						y:         nextY,
					})
					bestHP[nextX][nextY].initHP = initHP
					bestHP[nextX][nextY].currentHP = currentHP
				}
			}
		}

		queue = nextQueue
	}
	return bestHP[m-1][n-1].initHP
}

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	dungeon = [][]int{{0}}
	dungeon = [][]int{{0, 0, 0}, {-1, 0, 0}, {2, 0, -2}}
	fmt.Println(calculateMinimumHP(dungeon))
}
