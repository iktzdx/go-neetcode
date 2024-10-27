package palindromepartitioning_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/backtracking/palindrome-partitioning"
)

func Test_Partition(t *testing.T) {
	tests := map[string]struct {
		s    string
		want [][]string
	}{
		"case #1": {
			s:    "a",
			want: [][]string{{"a"}},
		},
		"case #2": {
			s:    "ab",
			want: [][]string{{"a", "b"}},
		},
		"case #3": {
			s:    "bb",
			want: [][]string{{"b", "b"}, {"bb"}},
		},
		"case #4": {
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		"case #5": {
			s: "aabaa",
			want: [][]string{
				{"a", "a", "b", "a", "a"},
				{"a", "a", "b", "aa"},
				{"aa", "b", "a", "a"},
				{"a", "aba", "a"},
				{"aa", "b", "aa"},
				{"aabaa"},
			},
		},
		"case #6": {
			s:    "abac",
			want: [][]string{{"a", "b", "a", "c"}, {"aba", "c"}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.Partition(test.s)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("Partition(%s): expected = %v, got = %v", test.s, test.want, got)
			}
		})
	}
}
