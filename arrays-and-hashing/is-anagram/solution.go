package validanagram

func IsAnagram(s string, t string) bool {
	// A slice of integers representing the amount of appearances
	// of each of the 26 lowercase English letters.
	chars := make([]int, 26)

	for _, v := range s {
		i := (v - 'a')
		chars[i]++
	}

	for _, v := range t {
		i := (v - 'a')
		chars[i]--
	}

	// If chars contains only zeros, then it's anagram.
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
