package longestrepeatingcharacterreplacement

func CharacterReplacement(s string, k int) int {
	var mostFreq, res int

	if len(s) < 2 || len(s) == k {
		return len(s)
	}

	freqs := make(map[byte]int, 26)

	for l, r := 0, 0; r < len(s); r++ {
		freqs[s[r]]++

		mostFreq = max(mostFreq, freqs[s[r]])
		if (r-l+1)-mostFreq > k {
			freqs[s[l]]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
