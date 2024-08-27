package validparentheses

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []string{}
	brackets := map[string]string{"}": "{", ")": "(", "]": "["}

	for _, val := range s {
		match, found := brackets[string(val)]
		if found && len(stack) == 0 {
			return false
		}

		if found && match == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, string(val))
	}

	return len(stack) == 0
}
