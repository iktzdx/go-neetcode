package evaluatereversepolishnotation_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/stack/evaluate-reverse-polish-notation"
)

func Test_EvalRPN(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		"single element": {
			input: []string{"18"},
			want:  18,
		},
		"simple addition - (a+b)": {
			input: []string{"3", "4", "+"},
			want:  7,
		},
		"simple subtraction - (a-b)": {
			input: []string{"4", "3", "-"},
			want:  1,
		},
		"simple multiplication - (a*b)": {
			input: []string{"4", "3", "*"},
			want:  12,
		},
		"simple division - (a/b)": {
			input: []string{"8", "2", "/"},
			want:  4,
		},
		"negative operands - (-a+-b)": {
			input: []string{"-2", "-4", "+"},
			want:  -6,
		},
		"multiple operators - ((a+b)*c)": {
			input: []string{"2", "1", "+", "3", "*"},
			want:  9,
		},
		"multiple operators - (a+(b/c))": {
			input: []string{"4", "13", "5", "/", "+"},
			want:  6,
		},
		"multiple operators - (a+b)*(c+d)": {
			input: []string{"3", "4", "+", "5", "6", "+", "*"},
			want:  77,
		},
		"long complex expression - (((a*(b/((c+d)*-x)))+y)+z)": {
			input: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:  22,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.EvalRPN(test.input)
			if test.want != got {
				t.Fatalf("EvalRPN(%v): expected = %v, got %v", test.input, test.want, got)
			}
		})
	}
}
