package anagramgroups

import "slices"

func GroupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}

	groups := make([][]string, 0)
	indexes := make(map[string]int, 0)

	for _, val := range strs {
		key := makeUniqueKey(val)

		if idx, found := indexes[key]; !found {
			indexes[key] = len(groups)
			groups = append(groups, []string{val})
		} else {
			groups[idx] = append(groups[idx], val)
		}
	}

	return groups
}

func makeUniqueKey(s string) string {
	key := []byte(s)

	slices.Sort(key)

	return string(key)
}
