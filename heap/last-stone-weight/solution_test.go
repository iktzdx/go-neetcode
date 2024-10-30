package laststoneweight_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/heap/last-stone-weight"
)

func Test_LastStoneWeight(t *testing.T) {
	tests := map[string]struct {
		stones []int
		want   int
	}{
		"case #1": {
			stones: []int{1},
			want:   1,
		},
		"case #2": {
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		"case #3": {
			stones: []int{2, 3, 6, 2, 4},
			want:   1,
		},
		"case #4": {
			stones: []int{1, 2},
			want:   1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.LastStoneWeight(test.stones)

			if test.want != got {
				t.Fatalf("LastStoneWeight(%v): expected = %d, got = %d", test.stones, test.want, got)
			}
		})
	}
}
