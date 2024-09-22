package lrucache_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/linked-list/lru-cache"
)

const evicted = -1

func Test_LRUCache(t *testing.T) {
	c := solution.Constructor(2)

	c.Put(1, 1) // LRUCache = [[1,1]]
	c.Put(2, 2) // LRUCache = [[1,1], [2,2]]

	require.Equal(t, 1, c.Get(1), "node [1, 1] must exist")

	// The number of keys exceeds the capacity,
	// evict the least recently used node - [2, 2].
	c.Put(3, 3) // LRUCache = [[1,1], [3,3]]

	require.Equal(t, evicted, c.Get(2), "node [2, 2] must be evicted")

	// The number of keys exceeds the capacity,
	// evict the least recently used node - [1, 1].
	c.Put(4, 4) // LRUCache = [[3,3], [4,4]]

	require.Equal(t, evicted, c.Get(1), "node [1, 1] must be evicted")

	require.Equal(t, 3, c.Get(3), "node [3, 3] must exist")
	require.Equal(t, 4, c.Get(4), "node [4, 4] must exist")

	// Update the value of the existing node.
	c.Put(4, 5) // LRUCache = [[3,3], [4,5]]
}
