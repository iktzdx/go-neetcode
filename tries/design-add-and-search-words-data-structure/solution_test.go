package designaddandsearchwordsdatastructure_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	solution "github.com/iktzdx/go-neetcode/tries/design-add-and-search-words-data-structure"
)

func Test_WordDictionary(t *testing.T) {
	dict := solution.Constructor()
	require.NotZero(t, dict, "create a new instance of WordDictionary")

	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")

	require.False(t, dict.Search("pad"))

	require.True(t, dict.Search("bad"))
	require.True(t, dict.Search(".ad"))
	require.True(t, dict.Search("b.."))

	require.False(t, dict.Search("..e"))

	dict.AddWord("apple")
	require.False(t, dict.Search("app"))
}
