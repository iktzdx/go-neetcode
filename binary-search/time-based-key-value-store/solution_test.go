package timebasedkeyvaluestore_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/binary-search/time-based-key-value-store"
)

func Test_TimeMap(t *testing.T) {
	timeMap := solution.Constructor()
	require.NotNil(t, timeMap, "construct time map")

	key1, val1, ts1 := "foo", "bar", 1
	timeMap.Set(key1, val1, ts1)
	require.Equalf(t, 1, timeMap.Len(), "set first value=%q to the key=%q at a timestamp=%d", val1, key1, ts1)

	require.Equalf(t, val1, timeMap.Get(key1, ts1), "get value=%q at a timestamp=%d from the key=%q", val1, ts1, key1)
	// Return "bar", since there is no value corresponding to foo at
	// timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	require.Equal(t, val1, timeMap.Get(key1, 2), "there is no corresponding timestamp, so get value=%q from the key=%q", val1, key1)

	// Store multiple values for the same key at different time stamps.
	val2, ts2 := "bar2", 4
	timeMap.Set(key1, val2, ts2)
	require.Equalf(t, 1, timeMap.Len(), "set second value=%q to the key=%q at a timestamp=%d", val2, key1, ts2)

	require.Equal(t, val2, timeMap.Get(key1, ts2), "get second value=%q at a timestamp=%d from the key=%q", val2, ts2, key1)
	// Return "bar2".
	require.Equal(t, val2, timeMap.Get(key1, 5), "there is no corresponding timestamp, so get value=%q from the key=%q", val2, key1)
	// If there are no values, it returns "".
	require.Zero(t, timeMap.Get("nonexistent", 6), "there is no such key, so return an empty string")
}
