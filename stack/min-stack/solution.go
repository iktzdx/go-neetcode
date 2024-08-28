package minstack

type MinStack struct {
	data []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		mins: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)

	minVal := val
	if len(s.mins) != 0 {
		minVal = min(s.mins[len(s.mins)-1], val)
	}

	s.mins = append(s.mins, minVal)
}

func (s *MinStack) Pop() {
	s.data = s.data[:len(s.data)-1]
	s.mins = s.mins[:len(s.mins)-1]
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}

func (s *MinStack) Len() int {
	return len(s.data)
}
