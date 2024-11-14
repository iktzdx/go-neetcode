package coursescheduleii_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/graphs/course-schedule-ii"
)

// func FindOrder(numCourses int, prerequisites [][]int) []int {

func Test_FindOrder(t *testing.T) {
	tests := map[string]struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		"case #1": {
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 1},
		},
		"case #2": {
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          []int{0, 2, 1, 3},
		},
		"case #3": {
			numCourses:    1,
			prerequisites: [][]int{},
			want:          []int{0},
		},
		"case #4": {
			numCourses:    3,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 1, 2},
		},
		"case #5": {
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			want:          []int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.FindOrder(test.numCourses, test.prerequisites)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("FindOrder(%d, %v): expected = %v, got = %v", test.numCourses, test.prerequisites, test.want, got)
			}
		})
	}
}
