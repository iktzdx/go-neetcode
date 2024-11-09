package wallsandgates_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/graphs/walls-and-gates"
)

func Test_WallsAndGates(t *testing.T) {
	tests := map[string]struct {
		rooms [][]int
		want  [][]int
	}{
		"case #1": {
			rooms: [][]int{},
			want:  [][]int{},
		},
		"case #2": {
			rooms: [][]int{
				{solution.Empty, solution.Wall, solution.Gate, solution.Empty},
				{solution.Empty, solution.Empty, solution.Empty, solution.Wall},
				{solution.Empty, solution.Wall, solution.Empty, solution.Wall},
				{solution.Gate, solution.Wall, solution.Empty, solution.Empty},
			},
			want: [][]int{
				{3, solution.Wall, solution.Gate, 1},
				{2, 2, 1, solution.Wall},
				{1, solution.Wall, 2, solution.Wall},
				{solution.Gate, solution.Wall, 3, 4},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			rooms := test.rooms[:]
			solution.WallsAndGates(rooms)

			if !reflect.DeepEqual(test.want, rooms) {
				t.Fatalf("WallsAndGates(%v): expected = %v, got = %v", test.rooms, test.want, rooms)
			}
		})
	}
}
