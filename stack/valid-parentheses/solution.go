package validparentheses

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}

	for _, val := range s {
		switch val {
		case '[', '(', '{':
			stack = append(stack, val)
		case ']', ')', '}':
			if len(stack) == 0 {
				return false
			}

			if !isMatching(stack[len(stack)-1], val) {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isMatching(o, c rune) bool {
	switch o {
	case '[':
		return c == ']'
	case '{':
		return c == '}'
	case '(':
		return c == ')'
	}

	return false
}
