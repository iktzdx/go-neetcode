package validanagram

func IsAnagram(s string, t string) bool {
	if s == t {
		return true
	}

	if len(s) != len(t) {
		return false
	}

	if compareMaps(countChars(s), countChars(t)) {
		return true
	}

	return false
}

func countChars(s string) map[rune]int {
	result := make(map[rune]int)

	for _, ch := range s {
		result[ch]++
	}

	return result
}

func compareMaps(x, y map[rune]int) bool {
	for k, xv := range x {
		yv, ok := y[k]
		if !ok {
			return false
		}

		if xv != yv {
			return false
		}
	}

	return true
}
