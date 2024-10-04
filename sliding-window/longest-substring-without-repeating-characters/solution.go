package longestsubstringwithoutrepeatingcharacters

func LengthOfLongestSubstring(s string) int {
	var result int

	if len(s) < 2 {
		return len(s)
	}

	chars := make(map[byte]bool, 0)

	for l, r := 0, 0; r < len(s); r++ {
		for ; chars[s[r]]; l++ {
			delete(chars, s[l])
		}

		chars[s[r]] = true
		result = max(result, r-l+1)
	}

	return result
}
