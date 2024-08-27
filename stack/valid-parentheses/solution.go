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
		if match, found := pairs[val]; found {
			if !stack.isEmpty() && stack.top() == match {
				stack.pop()
			} else {
				return false
			}
		} else {
			stack.push(val)
		}
	}

	return stack.isEmpty()
}
