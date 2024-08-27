package validparentheses

type stack []rune

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) push(val rune) {
	*s = append(*s, val)
}

func (s *stack) pop() rune {
	if s.isEmpty() {
		return -1
	}

	last := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]

	return last
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
		if found && stack.isEmpty() {
			return false
		}

		if !found {
			stack.push(val)
			continue
		}

		if last := stack.pop(); last != match {
			return false
		}
	}

	return stack.isEmpty()
}
