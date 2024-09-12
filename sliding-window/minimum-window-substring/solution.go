package minimumwindowsubstring

const maxLen = 100001

func MinWindow(s string, t string) string {
	if s == t {
		return t
	}

	need := make(map[rune]int, len(t))
	have := make(map[rune]int, 0)
	for _, ch := range t {
		need[ch]++
	}

	needCount, haveCount := len(t), 0
	res, resLen := [2]int{0, 0}, maxLen

	var result string

	left := 0
	for right := 0; right < len(s); right++ {
		ch := rune(s[right])
		have[ch]++

		if count, ok := need[ch]; ok && have[ch] == count {
			haveCount++
		}

		for haveCount == needCount {
			// update result
			if (right - left + 1) < resLen {
				res = [2]int{left, right}
				resLen = right - left + 1
			}

			// pop from the left
			ch := rune(s[left])
			have[ch]--

			if count, ok := need[ch]; ok && have[ch] < count {
				haveCount--
			}
			left++
		}
	}

	if resLen != maxLen {
		result = s[res[0] : res[1]+1]
	}

	return result
}
