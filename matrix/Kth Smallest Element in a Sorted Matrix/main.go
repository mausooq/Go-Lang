package main

import (
	"container/heap"
	"fmt"
)

type ele struct {
	val, i, j int
}

type minHeap []ele

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(ele))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	m := len(matrix[0])

	h := &minHeap{}
	heap.Init(h)

	// Push the first element of each row into the heap
	for i := 0; i < n; i++ {
		heap.Push(h, ele{matrix[i][0], i, 0})
	}

	// Pop the smallest element k-1 times
	for i := 0; i < k-1; i++ {
		curr := heap.Pop(h).(ele)
		 r, c := curr.i, curr.j
		if c+1 < m {
			heap.Push(h, ele{matrix[r][c+1], r, c + 1})
		}
	}

	// The top of the heap is the kth smallest element
	return heap.Pop(h).(ele).val
}

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k := 8

	fmt.Println("Kth Smallest Element:", kthSmallest(matrix, k)) // Output: 13
}
