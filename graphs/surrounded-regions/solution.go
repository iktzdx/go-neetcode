package surroundedregions

type coords struct {
	row, col int
}

type dfsFunc func(r, c int)

const (
	Captured     = 'X'
	NotCaptured  = 'O'
	Unsurrounded = 'T'
)

func SurroundedRegions(board [][]byte) {
	var capture dfsFunc

	rows, cols := len(board), len(board[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	capture = func(r, c int) {
		if r < 0 || c < 0 ||
			r == rows || c == cols ||
			board[r][c] != NotCaptured {
			return
		}

		board[r][c] = Unsurrounded

		for _, d := range directions {
			capture(r+d[0], c+d[1])
		}
	}

	// Capture unsurrounded regions ('O' -> 'T').
	for r := range rows {
		for c := range cols {
			if ((r == 0 || r == rows-1) || (c == 0 || c == cols-1)) &&
				board[r][c] == NotCaptured {
				capture(r, c)
			}
		}
	}

	// Capture surrounded regions ('O' -> 'X').
	for r := range rows {
		for c := range cols {
			if NotCaptured == board[r][c] {
				board[r][c] = Captured
			}
		}
	}

	// Uncapture unsurrounded regions ('T' -> 'O').
	for r := range rows {
		for c := range cols {
			if Unsurrounded == board[r][c] {
				board[r][c] = NotCaptured
			}
		}
	}
}
