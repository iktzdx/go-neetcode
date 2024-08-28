package minstack_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/stack/min-stack"
	"github.com/stretchr/testify/require"
)

func Test_MinStack(t *testing.T) {
	minStack := solution.Constructor()

	require.NotNil(t, minStack, "initialize empty stack")

	minStack.Push(1)
	minStack.Push(0)
	minStack.Push(-1)
	minStack.Push(2)
	minStack.Push(-1)

	require.Equal(t, minStack.Len(), 5, "stack size after several pushes")
	require.Equal(t, minStack.GetMin(), -1, "min stack element")
	require.Equal(t, minStack.Top(), -1, "top stack element")

	minStack.Pop()

	require.Equal(t, minStack.Len(), 4, "stack size after a single pop")
	require.Equal(t, minStack.Top(), 2, "new top stack element")
	require.Equal(t, minStack.GetMin(), -1, "new min stack element")
}
