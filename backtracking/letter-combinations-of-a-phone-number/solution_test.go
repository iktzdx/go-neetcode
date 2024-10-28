package lettercombinationsofaphonenumber_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/backtracking/letter-combinations-of-a-phone-number"
)

func Test_LetterCombinations(t *testing.T) {
	tests := map[string]struct {
		digits string
		want   []string
	}{
		"case #1": {
			digits: "",
			want:   []string{},
		},
		"case #2": {
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		"case #3": {
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		"case #4": {
			digits: "34",
			want:   []string{"dg", "dh", "di", "eg", "eh", "ei", "fg", "fh", "fi"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.LetterCombinations(test.digits)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("LetterCombinations(%s): expected = %v, got = %v", test.digits, test.want, got)
			}
		})
	}
}
