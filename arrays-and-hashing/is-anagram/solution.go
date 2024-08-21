package validanagram

import "sort"

func IsAnagram(s string, t string) bool {
	return sortStringByCharacters(s) == sortStringByCharacters(t)
}

func stringToRuneSlice(s string) []rune {
	res := make([]rune, len(s))

	for _, v := range s {
		res = append(res, v)
	}

	return res
}

func sortStringByCharacters(s string) string {
	r := stringToRuneSlice(s)

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}
