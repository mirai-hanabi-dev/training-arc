package main

import (
	"fmt"
	"sort"
)

type pole struct {
	isStart bool
	id      int
	value   int
	height  int
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	poles := make([]pole, 2*n)
	for i := 0; i < len(buildings); i++ {
		poles[i] = pole{
			isStart: true,
			id:      i,
			value:   buildings[i][0],
			height:  buildings[i][2],
		}

		poles[i+n] = pole{
			isStart: false,
			id:      i,
			value:   buildings[i][1],
			height:  buildings[i][2],
		}
	}

	sort.Slice(poles, func(i, j int) bool {
		if poles[i].value == poles[j].value {
			return poles[i].isStart
		}
		return poles[i].value < poles[j].value
	})

	skylines := [][]int{}
	for _, pole := range poles {
		if pole.isStart {

		} else {

		}
	}

	return skylines
}

func main() {
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}

	fmt.Println(getSkyline(buildings))
}
