package validanagram

func IsAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		if count, exists := chars[v]; !exists || count == 0 {
			return false
		}

		chars[v]--
	}

	return len(s) == len(t)
}
