package timebasedkeyvaluestore_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/binary-search/time-based-key-value-store"
)

const (
	nonExistentKey       = "nonexistent"
	nonExistentTimestamp = 69
)

func Test_TimeMap(t *testing.T) {
	timeMap := solution.Constructor()
	require.NotNil(t, timeMap, "construct time map")

	key1, val1, ts1 := "foo", "bar", 1
	timeMap.Set(key1, val1, ts1)
	require.Equalf(t, 1, timeMap.Len(),
		"set 1st value=%q to the key=%q at a timestamp=%d", val1, key1, ts1)

	require.Equalf(t, val1, timeMap.Get(key1, ts1),
		"get 1st value=%q at a timestamp=%d from the key=%q", val1, ts1, key1)
	// Return "bar", since there is no value corresponding to foo at
	// timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	require.Equal(t, val1, timeMap.Get(key1, nonExistentTimestamp),
		"there is no corresponding timestamp, get 1st value=%q from the key=%q", val1, key1)

	// Store multiple values for the same key at different time stamps.
	val2, ts2 := "bar2", 2
	timeMap.Set(key1, val2, ts2)
	require.Equalf(t, 1, timeMap.Len(),
		"set 2nd value=%q to the key=%q at a timestamp=%d", val2, key1, ts2)

	require.Equal(t, val2, timeMap.Get(key1, ts2),
		"get 2nd value=%q at a timestamp=%d from the key=%q", val2, ts2, key1)
	// Return "bar2".
	require.Equal(t, val2, timeMap.Get(key1, nonExistentTimestamp),
		"there is no corresponding timestamp, get 2nd value=%q from the key=%q", val2, key1)

	val3, ts3 := "bar3", 3
	timeMap.Set(key1, val3, ts3)
	require.Equalf(t, 1, timeMap.Len(),
		"set 3rd value=%q to the key=%q at a timestamp=%d", val3, key1, ts3)

	require.Equal(t, val1, timeMap.Get(key1, ts1),
		"get 1st value=%q at a timestamp=%d from the key=%q", val1, ts1, key1)

	// If there are no values, it returns "".
	require.Zero(t, timeMap.Get(nonExistentKey, nonExistentTimestamp),
		"there is no such key, return empty string")
}
