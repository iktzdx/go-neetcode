package minstack

type MinStack struct {
	data      []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{
		data:      make([]int, 0),
		minValues: make([]int, 0),
	}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)

	if len(s.minValues) == 0 || val <= s.minValues[len(s.minValues)-1] {
		s.minValues = append(s.minValues, val)
	}
}

func (s *MinStack) Pop() {
	last := s.data[s.Len()-1]

	if last <= s.minValues[len(s.minValues)-1] {
		s.minValues = s.minValues[:len(s.minValues)-1]
	}

	s.data = s.data[:s.Len()-1]
}

func (s *MinStack) Top() int {
	return s.data[s.Len()-1]
}

func (s *MinStack) GetMin() int {
	return s.minValues[len(s.minValues)-1]
}

func (s *MinStack) Len() int {
	return len(s.data)
}
