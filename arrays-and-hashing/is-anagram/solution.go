package validanagram

func IsAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		chars[v]--
	}

	// If chars contains only zeros, then it's anagram.
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
