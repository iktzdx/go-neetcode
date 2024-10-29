package nqueens_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/backtracking/n-queens"
)

func Test_SolveNQueens(t *testing.T) {
	tests := map[string]struct {
		n    int
		want [][]string
	}{
		"case #1": {
			n:    1,
			want: [][]string{{"Q"}},
		},
		"case #2": {
			n: 4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.SolveNQueens(test.n)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("SolveNQueens(%d): expected = %v, got= %v", test.n, test.want, got)
			}
		})
	}
}
