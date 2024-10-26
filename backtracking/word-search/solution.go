package wordsearch

type dfsFunc func(row, col, pos int) bool

func Exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var dfs dfsFunc

	dfs = func(row, col, pos int) bool {
		if pos == len(word) {
			return true
		}

		if row < 0 || col < 0 ||
			row >= rows || col >= cols ||
			word[pos] != board[row][col] ||
			board[row][col] != word[pos] {
			return false
		}

		state := board[row][col]
		board[row][col] = '.'

		result := (dfs(row+1, col, pos+1) ||
			dfs(row-1, col, pos+1) ||
			dfs(row, col+1, pos+1) ||
			dfs(row, col-1, pos+1))

		board[row][col] = state

		return result
	}

	for r := range rows {
		for c := range cols {
			if board[r][c] == word[0] && dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}
