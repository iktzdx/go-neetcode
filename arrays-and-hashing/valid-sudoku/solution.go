package validsudoku

const (
	boardSideLength int  = 9
	boxSideLength   int  = 3
	emptyCell       byte = '.'
)

func IsValidSudoku(board [][]byte) bool {
	var (
		rows  [boardSideLength][boardSideLength]bool
		cols  [boardSideLength][boardSideLength]bool
		boxes [boxSideLength][boxSideLength][boardSideLength]bool
	)

	for row := 0; row < boardSideLength; row++ {
		for col := 0; col < boardSideLength; col++ {
			if board[row][col] == emptyCell {
				continue
			}

			v := int(board[row][col]-'0') - 1

			if rows[row][v] || cols[col][v] || boxes[row/boxSideLength][col/boxSideLength][v] {
				return false
			}

			rows[row][v], cols[col][v], boxes[row/3][col/3][v] = true, true, true
		}
	}

	return true
}
