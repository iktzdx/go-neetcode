package validparentheses

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	brackets := map[rune]rune{'}': '{', ')': '(', ']': '['}

	for _, val := range s {
		match, found := brackets[val]
		if found && len(stack) == 0 {
			return false
		}

		if found && match == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, val)
	}

	return len(stack) == 0
}
