package generateparentheses

import (
	"strings"
)

// GenerateParenthesis generates all combinations of well-formed parentheses.
func GenerateParenthesis(n int) []string {
	var stack, res []string
	var backtracking func(open, closed int)

	backtracking = func(open, closed int) {
		// Base case, stack will contain the proper parentheses.
		if open == n && closed == n {
			// Append a valid combination as a string.
			res = append(res, strings.Join(stack, ""))
			return
		}

		// Only add open parenthesis if open counter < n.
		if open < n {
			stack = append(stack, "(")
			backtracking(open+1, closed)
			// After backtracking returns, we have to update (clean up) our stack
			// by popping a character that we just added.
			stack = stack[:len(stack)-1]
		}

		// Only add closing parenthesis if closed counter < open counter.
		if closed < open {
			stack = append(stack, ")")
			backtracking(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtracking(0, 0)

	return res
}
