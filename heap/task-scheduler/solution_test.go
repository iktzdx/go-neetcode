package taskscheduler_test

import (
	"testing"

	taskscheduler "github.com/iktzdx/go-neetcode/heap/task-scheduler"
)

func Test_LeastInterval(t *testing.T) {
	tests := map[string]struct {
		tasks []byte
		n     int
		want  int
	}{
		"case #1": {
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			want:  8, // A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
		},
		"case #2": {
			tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'},
			n:     1,
			want:  6, // A possible sequence is: A -> B -> C -> D -> A -> B.
		},
		"case #3": {
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     3,
			want:  10, // A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
		},
		"case #4": {
			tasks: []byte{'X', 'X', 'Y', 'Y'},
			n:     2,
			want:  5, // A possible sequence is: X -> Y -> idle -> X -> Y.
		},
		"case #5": {
			tasks: []byte{'A', 'A', 'A', 'B', 'C'},
			n:     3,
			want:  9, // A possible sequence is: A -> B -> C -> Idle -> A -> Idle -> Idle -> Idle -> A.
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := taskscheduler.LeastInterval(test.tasks, test.n)

			if test.want != got {
				t.Fatalf("LeastInterval(%v, %d): expected = %d, got = %v", test.tasks, test.n, test.want, got)
			}
		})
	}
}
