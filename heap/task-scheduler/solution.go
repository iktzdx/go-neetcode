package taskscheduler

import "container/heap"

const size = 26

func LeastInterval(tasks []byte, n int) int {
	// Collect count for each task.
	count := make([]int, size)
	for _, t := range tasks {
		count[t-'A']++
	}

	h := &taskHeap{}
	heap.Init(h)

	// Heapify.
	for _, freq := range count {
		if freq > 0 {
			heap.Push(h, freq)
		}
	}

	time := 0
	q := make([][2]int, 0)

	for h.Len() > 0 || len(q) > 0 {
		time++

		if h.Len() > 0 {
			freq, _ := heap.Pop(h).(int)

			if freq-1 > 0 {
				q = append(q, [2]int{freq - 1, time + n})
			}
		}

		if len(q) > 0 && q[0][1] == time {
			// Pop left.
			freq := q[0][0]
			q = q[1:]

			heap.Push(h, freq)
		}
	}

	return time
}

type taskHeap []int

func (h taskHeap) Len() int {
	return len(h)
}

func (h taskHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h taskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *taskHeap) Push(x interface{}) {
	val, _ := x.(int)

	*h = append(*h, val)
}

func (h *taskHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]

	return popped
}
