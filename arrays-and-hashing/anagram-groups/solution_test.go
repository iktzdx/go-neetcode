package anagramgroups_test

import (
	"reflect"
	"testing"

	anagramgroups "github.com/iktzdx/go-neetcode/arrays-and-hashing/anagram-groups"
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := anagramgroups.GroupAnagrams(test.strs)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("GroupAnagrams(%v): expected = %v, got = %v", test.strs, test.want, got)
			}
		})
	}
}
