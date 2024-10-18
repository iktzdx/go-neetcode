package implementtrieprefixtree_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/tries/implement-trie-prefix-tree"
)

func Test_Trie(t *testing.T) {
	obj := solution.Constructor()
	require.NotZero(t, obj, "create a new trie instance")

	obj.Insert("apple")
	require.True(t, obj.Search("apple"), "search for a previously inserted word")
	require.False(t, obj.Search("app"), "the last node is not marked as the end of the word")
	require.True(t, obj.StartsWith("app"), "there should be a word with such prefix")

	obj.Insert("app")
	require.True(t, obj.Search("app"), "search for a previously inserted word")

	require.False(t, obj.Search("ape"), "one of the characters does not exist in this trie")

	obj.Insert("ape")
	require.True(t, obj.Search("ape"), "search for a previously inserted word")

	require.False(t, obj.StartsWith("b"), "there should not be a word with such prefix")
}
