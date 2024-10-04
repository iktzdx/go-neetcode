package anagramgroups_test

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	solution "github.com/iktzdx/go-neetcode/arrays-and-hashing/anagram-groups"
)

func Test_GroupAnagrams(t *testing.T) {
	tests := map[string]struct {
		strs []string
		want [][]string
	}{
		"0 or more matches": {
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		"only one element": {
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		"empty string": {
			strs: []string{""},
			want: [][]string{{""}},
		},
		"with unicode characters": {
			strs: []string{"колба", "карат", "бокал", "карта", "монета", "катар"},
			want: [][]string{{"монета"}, {"карат", "карта", "катар"}, {"колба", "бокал"}},
		},
		"no anagrams": {
			strs: []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
			want: [][]string{{"max"}, {"buy"}, {"doc"}, {"may"}, {"ill"}, {"duh"}, {"tin"}, {"bar"}, {"pew"}, {"cab"}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.GroupAnagrams(test.strs)
			if !compareGroups(test.want, got) {
				t.Fatalf("GroupAnagrams(%v): expected = %v, got = %v", test.strs, test.want, got)
			}
		})
	}
}

func compareGroups(want, got [][]string) bool {
	if len(want) != len(got) {
		return false
	}

	return reflect.DeepEqual(sliceToMap(want), sliceToMap(got))
}

func sliceToMap(ss [][]string) map[string]int {
	m := make(map[string]int, 0)

	for _, s := range ss {
		slices.Sort(s)

		m[fmt.Sprint(s)]++
	}

	return m
}
