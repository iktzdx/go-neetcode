package numberofislands

type bfsFunc func(r, c int)

func NumIslands(grid [][]byte) int {
	var (
		islands int
		bfs     bfsFunc
	)

	if len(grid) == 0 {
		return islands
	}

	rows, cols := len(grid), len(grid[0])

	bfs = func(r, c int) {
		directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		q := [][2]int{{r, c}}

		grid[r][c] = '.'

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			for _, d := range directions {
				row, col := curr[0]+d[0], curr[1]+d[1]

				if (row >= 0 && row < rows) && (col >= 0 && col < cols) && grid[row][col] == '1' {
					q = append(q, [2]int{row, col})
					grid[row][col] = '.'
				}
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				islands++

				bfs(r, c)
			}
		}
	}

	return islands
}
