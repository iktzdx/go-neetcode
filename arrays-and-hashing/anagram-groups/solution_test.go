package anagramgroups_test

import (
	"slices"
	"sort"
	"testing"

	anagramgroups "github.com/iktzdx/go-neetcode/arrays-and-hashing/anagram-groups"
	"github.com/stretchr/testify/assert"
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
			got := anagramgroups.GroupAnagrams(test.strs)
			if !compareGroups(t, test.want, got) {
				t.Fatalf("GroupAnagrams(%v): expected = %v, got = %v", test.strs, test.want, got)
			}
		})
	}
}

func compareGroups(t *testing.T, want, got [][]string) bool {
	sort.Slice(want, func(i, j int) bool {
		if len(want[i]) == len(want[j]) {
			slices.Sort(want[i])
			slices.Sort(want[j])

			return want[i][0] < want[j][0]
		}

		return len(want[i]) < len(want[j])
	})

	sort.Slice(got, func(i, j int) bool {
		if len(got[i]) == len(got[j]) {
			slices.Sort(got[i])
			slices.Sort(got[j])

			return got[i][0] < got[j][0]
		}

		return len(got[i]) < len(got[j])
	})

	for i, group := range got {
		if !assert.ElementsMatch(t, want[i], group) {
			return false
		}
	}

	return true
}
