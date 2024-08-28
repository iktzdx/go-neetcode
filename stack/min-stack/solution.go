package minstack

import "slices"

type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{data: make([]int, 0)}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *MinStack) Pop() {
	s.data = s.data[:s.Len()-1]
}

func (s *MinStack) Top() int {
	return s.data[s.Len()-1]
}

func (s *MinStack) GetMin() int {
	return slices.Min(s.data)
}

func (s *MinStack) Len() int {
	return len(s.data)
}
