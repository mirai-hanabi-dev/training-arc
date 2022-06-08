package main

import (
	"container/heap"
)

// Keep track of (n+1)/2 smallest elements and largest elements.

type MedianFinder struct {
	size    int
	smaller maxHeap
	larger  minHeap
}

func Constructor() MedianFinder {
	obj := MedianFinder{
		smaller: maxHeap{},
		larger:  minHeap{},
	}
	return obj
}

func (this *MedianFinder) AddNum(num int) {
	this.size += 1
	if this.size == 1 {
		heap.Push(&this.smaller, num)
		heap.Push(&this.larger, num)
		return
	}
	heapExpectedSize := (this.size + 1) / 2
	if len(this.smaller) < heapExpectedSize { // This condition applies for both heaps
		if num < this.smaller[0] {
			heap.Push(&this.larger, this.smaller[0])
			heap.Push(&this.smaller, num)
		} else if num > this.larger[0] {
			heap.Push(&this.larger, num)
			heap.Push(&this.smaller, this.larger[0])
		} else {
			heap.Push(&this.smaller, num)
			heap.Push(&this.larger, num)
		}
	} else {
		if num < this.smaller[0] {
			heap.Pop(&this.smaller)
			heap.Push(&this.smaller, num)
		}
		if num > this.larger[0] {
			heap.Pop(&this.larger)
			heap.Push(&this.larger, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	return float64(this.smaller[0]+this.larger[0]) / 2
}

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
