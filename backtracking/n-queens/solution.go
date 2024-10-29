package nqueens

import "strings"

const (
	figure = "Q"
	empty  = "."
)

type (
	node struct {
		cols     map[int]bool
		posDiags map[int]bool // (row + col)
		negDiags map[int]bool // (row - col)
	}

	backtracking func(row int)
)

func SolveNQueens(n int) [][]string {
	result := make([][]string, 0)
	b := board(n)

	nn := node{
		cols:     make(map[int]bool),
		posDiags: make(map[int]bool),
		negDiags: make(map[int]bool),
	}

	var backtrack backtracking

	backtrack = func(row int) {
		if row == n {
			// [".", ".", ".", "."] -> ["...."]
			bCopy := make([]string, n)
			for i, row := range b {
				bCopy[i] = strings.Join(row, "")
			}

			result = append(result, bCopy)

			return
		}

		for col := range n {
			if nn.cols[col] || nn.posDiags[row+col] || nn.negDiags[row-col] {
				continue
			}

			nn.cols[col] = true
			nn.posDiags[row+col] = true
			nn.negDiags[row-col] = true

			b[row][col] = figure

			backtrack(row + 1)

			nn.cols[col] = false
			nn.posDiags[row+col] = false
			nn.negDiags[row-col] = false

			b[row][col] = empty
		}
	}

	backtrack(0)

	return result
}

func board(n int) [][]string {
	b := make([][]string, n)

	for i := range b {
		b[i] = make([]string, n)
		for j := range b[i] {
			b[i][j] = empty
		}
	}

	return b
}
