package nqueens

import (
	"strings"
)

const (
	figure = "Q"
	empty  = "."
)

type (
	node struct {
		cols    []bool
		posDiag []bool // row + col
		negDiag []bool // row - col
	}

	backtracking func(row int)
)

func SolveNQueens(n int) [][]string {
	result := make([][]string, 0)
	b := board(n)

	nn := node{
		cols:    make([]bool, n),
		posDiag: make([]bool, 2*n),
		negDiag: make([]bool, 2*n),
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
			// Check if valid placement.
			if nn.cols[col] || nn.posDiag[row+col] || nn.negDiag[row-col+n] {
				continue
			}

			// Place the Queen.
			nn.cols[col], nn.posDiag[row+col], nn.negDiag[row-col+n] = true, true, true
			b[row][col] = figure

			backtrack(row + 1)

			// Undo changes.
			nn.cols[col], nn.posDiag[row+col], nn.negDiag[row-col+n] = false, false, false
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
