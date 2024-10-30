package laststoneweight

import "container/heap"

func LastStoneWeight(stones []int) int {
	h := &maxHeap{}

	heap.Init(h)

	for _, s := range stones {
		heap.Push(h, s)
	}

	for h.Len() > 1 {
		s1, _ := heap.Pop(h).(int)
		s2, _ := heap.Pop(h).(int)

		if s1 != s2 {
			heap.Push(h, s1-s2)
		}
	}

	var result int
	if h.Len() == 1 {
		result = (*h)[0]
	}

	return result
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	val, _ := x.(int)

	*h = append(*h, val)
}

func (h *maxHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]

	return popped
}
