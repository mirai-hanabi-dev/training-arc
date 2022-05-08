// Optimal solution is O(n).
// However I come up with this fun O(nlogn) solution,
// then decide to stick with it.
package main

import (
	"container/heap"
	"fmt"
)

func candy(rates []int) int {
	pq := make(PriorityQueue, len(rates))
	ans := make([]int, len(rates))

	for idx, rate := range rates {
		pq[idx] = &Item{
			value: rate,
			index: idx,
		}
	}

	heap.Init(&pq)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		idx := item.index
		ans[idx] = 1
		if idx > 0 && rates[idx-1] != rates[idx] {
			ans[idx] = max(ans[idx], ans[idx-1]+1)
		}
		if idx < len(rates)-1 && rates[idx+1] != rates[idx] {
			ans[idx] = max(ans[idx], ans[idx+1]+1)
		}
	}

	fmt.Println(ans)
	res := 0
	for _, candy := range ans {
		res += candy
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Item struct {
	value int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].value < pq[j].value }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func main() {
	ratings := []int{9, 3, 4, 5, 7, 2, 1, 4, 4, 4, 2, 6, 6}
	fmt.Println(candy(ratings))
}
