package minstack_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/stack/min-stack"
)

func Test_MinStack(t *testing.T) {
	minStack := solution.Constructor()

	require.NotNil(t, minStack, "initialize empty stack")

	minStack.Push(1)
	minStack.Push(0)
	minStack.Push(-1)
	minStack.Push(2)
	minStack.Push(-1)

	require.Equal(t, 5, minStack.Len(), "stack size after several pushes")
	require.Equal(t, minStack.GetMin(), -1, "min stack element")
	require.Equal(t, minStack.Top(), -1, "top stack element")

	minStack.Pop()

	require.Equal(t, 4, minStack.Len(), "stack size after a single pop")
	require.Equal(t, 2, minStack.Top(), "new top stack element")
	require.Equal(t, minStack.GetMin(), -1, "new min stack element")
}
