package climbingstairs_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/1d-dynamic-programming/climbing-stairs"
)

func Test_ClimbStairs(t *testing.T) {
	tests := map[string]struct {
		n    int
		want int
	}{
		"case #1": {
			n:    2,
			want: 2,
		},
		"case #2": {
			n:    3,
			want: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.ClimbStairs(test.n)

			if test.want != got {
				t.Fatalf("ClimbStairs(%d): expected = %d, got = %d", test.n, test.want, got)
			}
		})
	}
}
