package kthlargestelementinastream

import "container/heap"

type KthLargest struct {
	size    int
	minHeap *minHeap
}

func Constructor(k int, nums []int) *KthLargest {
	h := &minHeap{}

	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return &KthLargest{
		size:    k,
		minHeap: h,
	}
}

func (obj *KthLargest) Add(val int) int {
	heap.Push(obj.minHeap, val)

	if obj.minHeap.Len() > obj.size {
		heap.Pop(obj.minHeap)
	}

	return obj.minHeap.min()
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(val any) {
	if v, ok := val.(int); ok {
		*h = append(*h, v)
	}
}

func (h *minHeap) Pop() any {
	var val int

	val, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]

	return val
}

func (h *minHeap) min() int {
	return (*h)[0]
}
