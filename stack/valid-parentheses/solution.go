package validparentheses

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	brackets := map[string]string{"}": "{", ")": "(", "]": "["}

	if _, found := brackets[string(s[0])]; found {
		return false
	}

	stack := []string{}

	for _, val := range s {
		match, found := brackets[string(val)]
		if found && match == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, string(val))
	}

	return len(stack) == 0
}
