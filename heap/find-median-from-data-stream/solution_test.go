package findmedianfromdatastream_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/heap/find-median-from-data-stream"
)

const delta float64 = 0.1

func Test_MediaFinder(t *testing.T) {
	finder := solution.Constructor()
	require.NotZero(t, finder, "create a new instance of MedianFinder")

	// [3], []
	finder.AddNum(3)
	require.InEpsilon(t, 3.0, finder.FindMedian(), delta, "keep added element in the left half of the stream")

	// [3, 2], []
	// len(left half) > len(right half) + threshold
	// [2], [3]
	finder.AddNum(2)
	require.InEpsilon(
		t, 2.5, finder.FindMedian(), delta,
		"after balancing, the maximum value from the left half is removed and added to the right half",
	)

	// [2], [3, 7]
	// len(left half) > len(right half) + threshold
	// [2, 3], [7]
	finder.AddNum(7)
	require.InEpsilon(
		t, 3, finder.FindMedian(), delta,
		"after balancing, the minimum value from the right half is removed and added to the left half",
	)

	// [2, 3], [7, 4]
	finder.AddNum(4)
	require.InEpsilon(t, 3.5, finder.FindMedian(), delta, "both halves are balanced and have equal sizes")
}
