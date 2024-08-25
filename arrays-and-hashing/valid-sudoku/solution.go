package validsudoku

import (
	"fmt"
)

const (
	boardSideLength int  = 9
	boxSideLength   int  = 3
	emptyCell       byte = '.'
)

type (
	rowsHashSet  map[int][]byte
	colsHashSet  map[int][]byte
	boxesHashSet map[string][]byte
)

func (h *boxesHashSet) add(k string, v byte) {
	set := *h
	set[k] = append(set[k], v)
	*h = set
}

func (h *boxesHashSet) contains(k string, target byte) bool {
	set := *h
	for _, val := range set[k] {
		if val == target {
			return true
		}
	}

	return false
}

func (h *rowsHashSet) add(k int, v byte) {
	set := *h
	set[k] = append(set[k], v)
	*h = set
}

func (h *rowsHashSet) contains(k int, target byte) bool {
	set := *h
	for _, val := range set[k] {
		if val == target {
			return true
		}
	}

	return false
}

func (h *colsHashSet) add(k int, v byte) {
	set := *h
	set[k] = append(set[k], v)
	*h = set
}

func (h *colsHashSet) contains(k int, target byte) bool {
	set := *h
	for _, val := range set[k] {
		if val == target {
			return true
		}
	}

	return false
}

func IsValidSudoku(board [][]byte) bool {
	rows := &rowsHashSet{}
	cols := &colsHashSet{}
	boxes := &boxesHashSet{}

	for row := 0; row < boardSideLength; row++ {
		for col := 0; col < boardSideLength; col++ {
			if board[row][col] == emptyCell {
				continue
			}

			if rows.contains(row, board[row][col]) {
				return false
			}

			if cols.contains(col, board[row][col]) {
				return false
			}

			if boxes.contains(hashKey(row, col), board[row][col]) {
				return false
			}

			rows.add(row, board[row][col])
			cols.add(col, board[row][col])
			boxes.add(hashKey(row, col), board[row][col])
		}
	}

	return true
}

func hashKey(row, col int) string {
	return fmt.Sprintf("%d,%d", row/boxSideLength, col/boxSideLength)
}
