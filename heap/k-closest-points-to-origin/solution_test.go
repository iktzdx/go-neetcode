package kclosestpointstoorigin_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	kclosestpointstoorigin "github.com/iktzdx/go-neetcode/heap/k-closest-points-to-origin"
)

func Test_KClosest(t *testing.T) {
	tests := map[string]struct {
		points [][]int
		k      int
		want   [][]int
	}{
		"case #1": {
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
		"case #2": {
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		"case #3": {
			points: [][]int{{0, 2}, {2, 2}},
			k:      1,
			want:   [][]int{{0, 2}},
		},
		"case #4": {
			points: [][]int{{0, 2}, {2, 0}, {2, 2}},
			k:      2,
			want:   [][]int{{0, 2}, {2, 0}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := kclosestpointstoorigin.KClosest(test.points, test.k)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("KClosest(%v, %d): expected = %v, got = %v", test.points, test.k, test.want, got)
			}
		})
	}
}
