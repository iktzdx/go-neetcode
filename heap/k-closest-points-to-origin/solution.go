package kclosestpointstoorigin

import (
	"container/heap"
)

type point struct {
	dist int
	x    int
	y    int
}

func KClosest(points [][]int, k int) [][]int {
	h := &pointHeap{}

	heap.Init(h)

	for _, pt := range points {
		x, y := pt[0], pt[1]
		dist := (x * x) + (y * y)

		heap.Push(h, point{dist, x, y})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([][]int, 0)

	for k > 0 {
		p, _ := heap.Pop(h).(point)

		result = append(result, []int{p.x, p.y})

		k--
	}

	return result
}

type pointHeap []point

func (h pointHeap) Len() int {
	return len(h)
}

func (h pointHeap) Less(i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h pointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pointHeap) Push(x interface{}) {
	val, _ := x.(point)

	*h = append(*h, val)
}

func (h *pointHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]

	return popped
}
