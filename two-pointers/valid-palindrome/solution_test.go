package validpalindrome_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/two-pointers/valid-palindrome"
)

func Test_IsPalindrome(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"is valid palindrome": {
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		"is invalid palindrome": {
			input: "race a car",
			want:  false,
		},
		"empty string": {
			input: " ",
			want:  true,
		},
		"all non-alphanumeric": {
			input: ",.",
			want:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.IsPalindrome(test.input)
			if test.want != got {
				t.Fatalf("IsPalindrome(%v): expected = %v, got = %v", test.input, test.want, got)
			}
		})
	}
}
