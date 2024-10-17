package kthlargestelementinastream_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/heap/kth-largest-element-in-a-stream"
)

func Test_KthLargest(t *testing.T) {
	obj := solution.Constructor(3, []int{4, 5, 8, 2})
	require.NotZero(t, obj, "new instance of KthLargest")

	require.Equal(t, 4, obj.Add(3))
	require.Equal(t, 5, obj.Add(5))
	require.Equal(t, 5, obj.Add(10))
	require.Equal(t, 8, obj.Add(9))
	require.Equal(t, 8, obj.Add(4))
}
