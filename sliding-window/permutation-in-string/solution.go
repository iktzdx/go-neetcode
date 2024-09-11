package permutationinstring

import "reflect"

func CheckInclusion(s1 string, s2 string) bool {
	s1Count := make(map[byte]int, len(s1))
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]]++
	}

	s2Count := make(map[byte]int, len(s1))

	left := 0
	for right := left + len(s1); right <= len(s2); right++ {
		currWindow := s2[left:right]
		for i := 0; i < len(currWindow); i++ {
			s2Count[currWindow[i]]++
		}

		if reflect.DeepEqual(s1Count, s2Count) {
			return true
		}

		clear(s2Count)
		left++
	}

	return false
}
