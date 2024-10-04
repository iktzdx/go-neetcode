package permutationinstring

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count := [26]int{}
	for _, ch := range s1 {
		count[ch-'a']++
	}

	left := 0

	for right := range s2 {
		count[s2[right]-'a']--

		if count == [26]int{} {
			return true
		}

		if right+1 >= len(s1) {
			count[s2[left]-'a']++
			left++
		}
	}

	return false
}
