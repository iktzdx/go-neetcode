package palindromepartitioning

import (
	validpalindrome "github.com/iktzdx/go-neetcode/two-pointers/valid-palindrome"
)

type backtracking func(pos int)

func Partition(s string) [][]string {
	result := make([][]string, 0)
	curr := make([]string, 0)

	var fn backtracking

	fn = func(pos int) {
		if pos >= len(s) {
			currCopy := make([]string, len(curr))
			copy(currCopy, curr)

			result = append(result, currCopy)

			return
		}

		for i := pos; i < len(s); i++ {
			substr := s[pos : i+1]

			if validpalindrome.IsPalindrome(substr) {
				curr = append(curr, substr)

				fn(i + 1)

				curr = curr[:len(curr)-1]
			}
		}
	}

	fn(0)

	return result
}
