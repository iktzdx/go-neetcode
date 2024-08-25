package validsudoku

const (
	boardSideLength int  = 9
	boxSideLength   int  = 3
	emptyCell       byte = '.'
)

func IsValidSudoku(board [][]byte) bool {
	return CheckRows(board) && CheckColumns(board) && CheckBoxes(board)
}

func CheckBoxes(board [][]byte) bool {
	boxMap := make(map[byte]struct{}, boardSideLength)

	var rowIdx, colIdx int
	for colIdx <= boardSideLength {
		if colIdx%boxSideLength == 0 && colIdx != 0 {
			rowIdx++
			colIdx -= boxSideLength

		}

		if (rowIdx%boxSideLength == 0 && rowIdx != 0) && (colIdx%boxSideLength == 0) {
			clear(boxMap)
		}

		if rowIdx == boardSideLength {
			rowIdx -= boardSideLength
			colIdx += boxSideLength
		}

		var boxVal byte
		if colIdx == boardSideLength {
			boxVal = board[rowIdx][colIdx-1]
		} else {
			boxVal = board[rowIdx][colIdx]
		}

		if _, exists := boxMap[boxVal]; exists && boxVal != emptyCell {
			return false
		}

		boxMap[boxVal] = struct{}{}

		colIdx++
	}

	return true
}

func CheckRows(board [][]byte) bool {
	rowMap := make(map[byte]struct{}, boardSideLength)

	for rowIdx := 0; rowIdx < boardSideLength; rowIdx++ {
		for colIdx := 0; colIdx < boardSideLength; colIdx++ {
			rowVal := board[rowIdx][colIdx]
			if _, exists := rowMap[rowVal]; exists && rowVal != emptyCell {
				return false
			}

			rowMap[rowVal] = struct{}{}

		}

		clear(rowMap)
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
