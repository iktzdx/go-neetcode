package kthlargestelementinanarray

import "container/heap"

func FindKthLargest(nums []int, k int) int {
	h := &numsHeap{}

	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
	}

	var result int
	for k > 0 {
		result, _ = heap.Pop(h).(int)

		k--
	}

	return result
}

type numsHeap []int

func (h numsHeap) Len() int {
	return len(h)
}

func (h numsHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h numsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *numsHeap) Push(x interface{}) {
	val, _ := x.(int)

	*h = append(*h, val)
}

func (h *numsHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]

	return popped
}
