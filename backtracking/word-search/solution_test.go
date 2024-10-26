package wordsearch_test

import (
	"testing"

	wordsearch "github.com/iktzdx/go-neetcode/backtracking/word-search"
)

func Test_WordSearch(t *testing.T) {
	tests := map[string]struct {
		board [][]byte
		word  string
		want  bool
	}{
		"case #1": {
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		"case #2": {
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		"case #3": {
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
		"case #4": {
			board: [][]byte{
				{'A', 'B', 'C', 'D'},
				{'S', 'A', 'A', 'T'},
				{'A', 'C', 'A', 'E'},
			},
			word: "CAT",
			want: true,
		},
		"case #5": {
			board: [][]byte{
				{'A', 'B', 'C', 'D'},
				{'S', 'A', 'A', 'T'},
				{'A', 'C', 'A', 'E'},
			},
			word: "BAT",
			want: false,
		},
		"case #6": {
			board: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			},
			word: "AAAAAAAAAAAAAAa",
			want: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := wordsearch.Exist(test.board, test.word)

			if test.want != got {
				t.Fatalf("Exist(%v, %v): expected = %v, got = %v", test.board, test.word, test.want, got)
			}
		})
	}
}
