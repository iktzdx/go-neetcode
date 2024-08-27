package validparentheses

type stack []rune

func (s *stack) push(val rune) {
	*s = append(*s, val)
}

func (s *stack) pop() {
	if !s.isEmpty() {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *stack) top() rune {
	if s.isEmpty() {
		return -1
	}

	return (*s)[len(*s)-1]
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := &stack{}
	pairs := map[rune]rune{']': '[', ')': '(', '}': '{'}

	for _, val := range s {
		match, found := pairs[val]
		switch {
		case found && stack.isEmpty():
			return false
		case found && stack.top() == match:
			stack.pop()
		default:
			stack.push(val)
		}
	}

	return stack.isEmpty()
}
