package copylistwithrandompointer_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/copy-list-with-random-pointer"
)

const null int = -1

func Test_CopyRandomList(t *testing.T) {
	tests := map[string]struct {
		head [][]int
		want [][]int
	}{
		"example #1": {
			head: [][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
			want: [][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		"example #2": {
			head: [][]int{{1, 1}, {2, 1}},
			want: [][]int{{1, 1}, {2, 1}},
		},
		"example #3": {
			head: [][]int{{3, null}, {3, 0}, {3, null}},
			want: [][]int{{3, null}, {3, 0}, {3, null}},
		},
		"example #4": {
			head: [][]int{},
			want: [][]int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			head := linkedListRandomFromArray(test.head)
			got := linkedListRandomToArray(solution.CopyRandomList(head))

			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("CopyRandomList(%v): expected = %v, got = %v", test.head, test.want, got)
			}
		})
	}
}

func linkedListRandomToArray(head *solution.Node) [][]int {
	result := [][]int{}
	curr := head
	nodes := map[*solution.Node]int{}
	i := 0

	for curr != nil {
		nodes[curr] = i

		result = append(result, []int{curr.Val, null})
		curr = curr.Next
		i++
	}

	// Set up random
	curr = head
	for i = 0; i < len(result); i++ {
		if curr.Random != nil {
			result[i][1] = nodes[curr.Random]
		}

		curr = curr.Next
	}

	return result
}

func linkedListRandomFromArray(arr [][]int) *solution.Node {
	if len(arr) == 0 {
		return nil
	}

	head := &solution.Node{
		Val: arr[0][0],
	}

	nodes := map[int]*solution.Node{0: head}

	// Build the linked list
	curr := head

	for i := 1; i < len(arr); i++ {
		next := &solution.Node{Val: arr[i][0]}
		nodes[i] = next
		curr.Next = next
		curr = next
	}

	// set up randoms
	curr = head

	for i := range arr {
		if arr[i][1] != null {
			if randNode, ok := nodes[arr[i][1]]; ok {
				curr.Random = randNode
			}
		}

		curr = curr.Next
	}

	return head
}
