package anagramgroups

import "slices"

func GroupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}

	storage := make(map[string][]string, 0)

	for _, val := range strs {
		key := makeUniqueKey(val)

		group := storage[key]
		if len(group) == 0 {
			storage[key] = []string{val}
		} else {
			storage[key] = append(group, val)
		}
	}

	idx := 0
	groups := make([][]string, len(storage))
	for _, group := range storage {
		groups[idx] = group
		idx++
	}

	return groups
}

func makeUniqueKey(s string) string {
	keyBytes := []byte(s)

	slices.Sort(keyBytes)

	return string(keyBytes)
}
