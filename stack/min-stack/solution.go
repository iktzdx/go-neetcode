package minstack

type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) isEmpty() bool {
	return s.len() == 0
}

func (s *stack) pop() {
	if !s.isEmpty() {
		*s = (*s)[:s.len()-1]
	}
}

func (s *stack) top() int {
	return (*s)[s.len()-1]
}

type MinStack struct {
	data      stack
	minValues stack
}

func Constructor() MinStack {
	return MinStack{
		data:      make(stack, 0),
		minValues: make(stack, 0),
	}
}

func (s *MinStack) Push(val int) {
	s.data.push(val)

	if s.minValues.isEmpty() || val <= s.minValues.top() {
		s.minValues.push(val)
	}
}

func (s *MinStack) Pop() {
	if s.data.top() <= s.minValues.top() {
		s.minValues.pop()
	}

	s.data.pop()
}

func (s *MinStack) Top() int {
	return s.data.top()
}

func (s *MinStack) GetMin() int {
	return s.minValues.top()
}

func (s *MinStack) Len() int {
	return s.data.len()
}
