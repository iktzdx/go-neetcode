package linkedlistcycle_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/linked-list/linked-list-cycle"
)

func Test_HasCycle_Example_1(t *testing.T) {
	node2 := &solution.ListNode{Val: 2}
	node3 := &solution.ListNode{Val: 3}
	node4 := &solution.ListNode{Val: 4}

	head := &solution.ListNode{Val: 1, Next: node2}

	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	assert.True(t, solution.HasCycle(head))
}

func Test_HasCycle_Example_2(t *testing.T) {
	node2 := &solution.ListNode{Val: 2}

	head := &solution.ListNode{Val: 1, Next: node2}

	node2.Next = head

	assert.True(t, solution.HasCycle(head))
}

func Test_HasCycle_Example_3(t *testing.T) {
	head := &solution.ListNode{Val: 1}

	assert.False(t, solution.HasCycle(head))
}
