package longestsubstringwithoutrepeatingcharacters

import (
	"slices"
)

func LengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var result int
	chars := make([]rune, 0)

	for _, ch := range s {
		for slices.Contains(chars, ch) {
			chars = chars[1:]
		}

		chars = append(chars, ch)
		result = max(result, len(chars))
	}

	return result
}
