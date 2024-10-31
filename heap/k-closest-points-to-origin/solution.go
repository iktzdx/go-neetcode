package kclosestpointstoorigin

import (
	"container/heap"
)

func KClosest(points [][]int, k int) [][]int {
	h := &minHeap{}

	heap.Init(h)

	for _, pt := range points {
		x, y := pt[0], pt[1]

		dist := (x * x) + (y * y)
		heap.Push(h, []int{dist, x, y})
	}

	result := make([][]int, 0)

	for k > 0 {
		val := heap.Pop(h)

		pts, _ := val.([]int)

		result = append(result, []int{pts[1], pts[2]})
		k--
	}

	return result
}

type minHeap [][]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	val, _ := x.([]int)

	*h = append(*h, val)
}

func (h *minHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]

	return popped
}
