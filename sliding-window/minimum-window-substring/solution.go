package minimumwindowsubstring

const maxLen = 100001

func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	} else if s == t {
		return t
	}

	var result string

	need := make(map[rune]int, len(t))
	for _, ch := range t {
		need[ch]++
	}

	needCount := len(t)
	start, resLen := 0, maxLen

	left, right := 0, 0
	for right < len(s) {
		ch := rune(s[right])
		if need[ch] > 0 {
			needCount--
		}

		need[ch]--
		right++

		for needCount == 0 {
			if right-left < resLen {
				start = left
				resLen = right - left
			}

			ch := rune(s[left])
			if need[ch] == 0 {
				needCount++
			}

			need[ch]++
			left++
		}
	}

	if resLen != maxLen {
		result = s[start : start+resLen]
	}

	return result
}
