package surroundedregions_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/graphs/surrounded-regions"
)

func Test_SurroundedRegions(t *testing.T) {
	tests := map[string]struct {
		board [][]byte
		want  [][]byte
	}{
		"case #1": {
			board: [][]byte{{solution.Captured}},
			want:  [][]byte{{solution.Captured}},
		},
		"case #2": {
			board: [][]byte{
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.NotCaptured, solution.NotCaptured, solution.Captured},
				{solution.Captured, solution.Captured, solution.NotCaptured, solution.Captured},
				{solution.Captured, solution.NotCaptured, solution.Captured, solution.Captured},
			},
			want: [][]byte{
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.NotCaptured, solution.Captured, solution.Captured},
			},
		},
		"case #3": {
			board: [][]byte{
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.NotCaptured, solution.NotCaptured, solution.Captured},
				{solution.Captured, solution.NotCaptured, solution.NotCaptured, solution.Captured},
				{solution.Captured, solution.Captured, solution.Captured, solution.NotCaptured},
			},
			want: [][]byte{
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.Captured, solution.Captured, solution.Captured},
				{solution.Captured, solution.Captured, solution.Captured, solution.NotCaptured},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			input := test.board

			solution.SurroundedRegions(input)

			output := test.board

			if !reflect.DeepEqual(test.want, output) {
				t.Fatalf("SurroundedRegions(%v): expected = %v, got = %v", input, test.want, output)
			}
		})
	}
}
