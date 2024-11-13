package courseschedule_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/graphs/course-schedule"
)

func Test_CanFinish(t *testing.T) {
	tests := map[string]struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		"case #1": {
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		"case #2": {
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		"case #3": {
			numCourses:    5,
			prerequisites: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4}},
			want:          true,
		},
		"case #4": {
			numCourses:    5,
			prerequisites: [][]int{{1, 2}, {3, 4}},
			want:          true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.CanFinish(test.numCourses, test.prerequisites)

			if test.want != got {
				t.Fatalf("CanFinish(%d, %v): expected = %v, got = %v", test.numCourses, test.prerequisites, test.want, got)
			}
		})
	}
}
