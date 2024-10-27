package palindromepartitioning

type backtracking func(pos int)

func Partition(s string) [][]string {
	result := make([][]string, 0)
	curr := make([]string, 0)

	var backtrack backtracking

	backtrack = func(pos int) {
		if pos >= len(s) {
			currCopy := make([]string, len(curr))
			copy(currCopy, curr)

			result = append(result, currCopy)

			return
		}

		for i := pos; i < len(s); i++ {
			if isValidPalindrome(s, pos, i) {
				curr = append(curr, s[pos:i+1])

				backtrack(i + 1)

				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack(0)

	return result
}

func isValidPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}
