package main

import (
	"container/heap"
	"sort"
)

type pole struct {
	isStart bool
	id      int
	value   int
	height  int
	poleEnd *pole
	item    *line
}

type line struct {
	id     int
	height int
	index  int
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	poles := make([]*pole, 2*n)
	for i := 0; i < len(buildings); i++ {
		poleEnd := &pole{
			isStart: false,
			id:      i,
			value:   buildings[i][1],
			height:  buildings[i][2],
		}
		poles[i+n] = poleEnd

		poles[i] = &pole{
			isStart: true,
			id:      i,
			value:   buildings[i][0],
			height:  buildings[i][2],
			poleEnd: poleEnd,
		}
	}

	sort.Slice(poles, func(i, j int) bool {
		if poles[i].value == poles[j].value {
			if poles[i].isStart && poles[j].isStart {
				return poles[i].height > poles[j].height
			}
			if !poles[i].isStart && !poles[j].isStart {
				return poles[i].height < poles[j].height
			}
			return poles[i].isStart
		}
		return poles[i].value < poles[j].value
	})

	skylines := [][]int{}

	pq := make(priorityqueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &line{id: -1, height: 0})

	for _, pole := range poles {
		top := pq[0]
		if pole.isStart {
			if top.height < pole.height {
				skylines = append(skylines, []int{pole.value, pole.height})
			}
			line := &line{id: pole.id, height: pole.height}
			pole.poleEnd.item = line
			heap.Push(&pq, line)
		} else {
			heap.Remove(&pq, pole.item.index)
			top := pq[0]
			if top.height < pole.height {
				skylines = append(skylines, []int{pole.value, top.height})
			}
		}
	}

	return skylines
}

type priorityqueue []*line

func (pq priorityqueue) Len() int { return len(pq) }

func (pq priorityqueue) Less(i, j int) bool {
	return pq[i].height > pq[j].height
}

func (pq priorityqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityqueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*line)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
