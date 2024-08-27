package validparentheses_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/stack/valid-parentheses"
)

func Test_IsValid(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid round brackets": {
			input: "()",
			want:  true,
		},
		"valid curly braces": {
			input: "{}",
			want:  true,
		},
		"valid square brackets": {
			input: "[]",
			want:  true,
		},
		"valid sequence of parentheses": {
			input: "()[]{}",
			want:  true,
		},
		"valid nested parentheses": {
			input: "({[]})",
			want:  true,
		},
		"invalid pair": {
			input: "(]",
			want:  false,
		},
		"closed in the wrong order": {
			input: "([{)]}",
			want:  false,
		},
		"invalid - no open brackets": {
			input: "))}}]]",
			want:  false,
		},
		"invalid - no closed brackets": {
			input: "(({{[[",
			want:  false,
		},
		"invalid - the opposite order": {
			input: ")(",
			want:  false,
		},
		"invalid - odd number of brackets": {
			input: "(({{[[]]}))",
			want:  false,
		},
		"invalid - wrong order in the end": {
			input: "(){}}{",
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.IsValid(test.input)
			if test.want != got {
				t.Fatalf("IsValid(%s): expect = %v, got = %v", test.input, test.want, got)
			}
		})
	}
}
