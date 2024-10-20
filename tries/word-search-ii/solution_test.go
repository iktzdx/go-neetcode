package wordsearchii_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/tries/word-search-ii"
)

func Test_FindWords(t *testing.T) {
	tests := map[string]struct {
		board [][]byte
		words []string
		want  []string
	}{
		"case #1": {
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words: []string{"abcb"},
			want:  []string{},
		},
		"case #2": {
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"eat", "oath"},
		},
		"case #3": {
			board: [][]byte{
				{'a', 'b', 'c', 'd'},
				{'s', 'a', 'a', 't'},
				{'a', 'c', 'k', 'e'},
				{'a', 'c', 'd', 'n'},
			},
			words: []string{"bat", "cat", "back", "backend", "stack"},
			want:  []string{"cat", "back", "backend"},
		},
		"case #4": {
			board: [][]byte{
				{'x', 'o'},
				{'x', 'o'},
			},
			words: []string{"xoxo"},
			want:  []string{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.FindWords(test.board, test.words)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("FindWords(%v, %v): expected = %v, got = %v", test.board, test.words, test.want, got)
			}
		})
	}
}
