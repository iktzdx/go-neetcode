package generateparentheses

import (
	"strings"
)

// GenerateParenthesis generates all combinations of well-formed parentheses.
func GenerateParenthesis(n int) []string {
	var (
		stack, res []string
		backtrack  func(int, int)
	)

	backtrack = func(open, closed int) {
		// Base case, stack will contain the proper parentheses.
		if open == n && closed == n && open == closed {
			// Append a valid combination as a string.
			res = append(res, strings.Join(stack, ""))

			return
		}

		// Only add open parenthesis if open counter < n.
		if open < n {
			stack = append(stack, "(")

			backtrack(open+1, closed)
			// After backtracking returns, we have to update (clean up) our stack
			// by popping a character that we just added.
			stack = stack[:len(stack)-1]
		}

		// Only add closing parenthesis if closed counter < open counter.
		if closed < open {
			stack = append(stack, ")")

			backtrack(open, closed+1)

			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return res
}
