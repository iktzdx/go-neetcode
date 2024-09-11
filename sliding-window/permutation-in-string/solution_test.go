package permutationinstring_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/sliding-window/permutation-in-string"
)

func Test_CheckInclusion(t *testing.T) {
	tests := map[string]struct {
		s1   string
		s2   string
		want bool
	}{
		"is a permutation": {
			s1:   "abc",
			s2:   "lecabee", // "cab"
			want: true,
		},
		"is a permutation itself": {
			s1:   "abc",
			s2:   "leabcee", // "abc"
			want: true,
		},
		"is not a permutation": {
			s1:   "abc",
			s2:   "lecaabee",
			want: false,
		},
		"s1 is a single character": {
			s1:   "x",
			s2:   "eidboaoo",
			want: false,
		},
		"same strings": {
			s1:   "abc",
			s2:   "abc",
			want: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.CheckInclusion(test.s1, test.s2)
			if test.want != got {
				t.Fatalf("CheckInclusion(%s, %s): expect = %v, got = %v", test.s1, test.s2, test.want, got)
			}
		})
	}
}
