package permutationinstring

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	got, want := [26]int{}, [26]int{}
	for _, ch := range s1 {
		want[ch-'a']++
	}

	for right := range s2 {
		if right >= len(s1) {
			got[s2[right-len(s1)]-'a']--
		}

		got[s2[right]-'a']++
		if got == want {
			return true
		}

	}

	return false
}
