package generateparentheses_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/stack/generate-parentheses"
	"github.com/stretchr/testify/assert"
)

func Test_GenerateParenthesis(t *testing.T) {
	tests := map[string]struct {
		input int
		want  []string
	}{
		"one pair": {
			input: 1,
			want:  []string{"()"},
		},
		"multiple pairs": {
			input: 3,
			want:  []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.GenerateParenthesis(test.input)
			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("GenerateParenthesis(%v): expected = %v, got = %v", test.input, test.want, got)
			}
		})
	}
}
