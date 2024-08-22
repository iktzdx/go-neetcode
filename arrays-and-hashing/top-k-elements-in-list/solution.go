package topkelementsinlist

import (
	"container/heap"
)

type (
	Item struct {
		key   int
		value int
	}

	MinHeap []*Item
)

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func TopKFrequent(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}

	freqs := make(map[int]int, k)
	for _, n := range nums {
		freqs[n]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for k, v := range freqs {
		heap.Push(h, &Item{key: k, value: v})
	}

	for h.Len() > k {
		heap.Pop(h)
	}

	res := make([]int, 0)
	for _, el := range *h {
		res = append(res, el.key)
	}

	return res
}
