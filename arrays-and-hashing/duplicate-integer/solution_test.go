package duplicateinteger_test

import (
	"testing"

	duplicateinteger "github.com/iktzdx/go-neetcode/arrays-and-hashing/duplicate-integer"
)

func Test_FindDuplicate(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"example 1": {
			input: []int{1, 3, 4, 2, 2},
			want:  2,
		},
		"example 2": {
			input: []int{3, 1, 3, 4, 2},
			want:  3,
		},
		"example 3": {
			input: []int{3, 3, 3, 3, 3},
			want:  3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := duplicateinteger.FindDuplicate(test.input)
			if test.want != got {
				t.Fatalf("FindDuplicate(%v), expected = %d, got = %d", test.input, test.want, got)
			}
		})
	}
}
