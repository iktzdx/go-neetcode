package validsudoku

const (
	boardSideLength int  = 9
	boxSideLength   int  = 3
	emptyCell       byte = '.'
)

func IsValidSudoku(board [][]byte) bool {
	for _, row := range board {
		if !CheckRow(row) {
			return false
		}
	}

	if !CheckColumns(board) {
		return false
	}

	if !CheckBoxes(board) {
		return false
	}

	return true
}

func CheckBoxes(board [][]byte) bool {
	return false
}

func CheckRow(row []byte) bool {
	rowMap := make(map[byte]struct{}, boardSideLength)

	for _, rowVal := range row {
		if _, exists := rowMap[rowVal]; exists && rowVal != emptyCell {
			return false
		}

		rowMap[rowVal] = struct{}{}
	}

	return true
}

func CheckColumns(board [][]byte) bool {
	colMap := make(map[byte]struct{}, boardSideLength)

	for colIdx := 0; colIdx < boardSideLength; colIdx++ {
		for rowIdx := 0; rowIdx < boardSideLength; rowIdx++ {
			colVal := board[rowIdx][colIdx]
			if _, exists := colMap[colVal]; exists && colVal != emptyCell {
				return false
			}

			colMap[colVal] = struct{}{}
		}
		clear(colMap)
	}

	return true
}
