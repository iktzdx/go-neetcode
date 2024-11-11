package pacificatlanticwaterflow

type dfsFunc func(r, c int, visit map[position]bool, prevHeight int)

type position struct {
	row, col int
}

func PacificAtlantic(heights [][]int) [][]int {
	var dfs dfsFunc

	rows, cols := len(heights), len(heights[0])
	directions := []position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	pac := make(map[position]bool)
	atl := make(map[position]bool)

	result := make([][]int, 0)

	dfs = func(r, c int, visit map[position]bool, prevHeight int) {
		pos := position{r, c}

		if visit[pos] ||
			r < 0 || r == rows ||
			c < 0 || c == cols ||
			heights[r][c] < prevHeight {
			return
		}

		visit[pos] = true

		for _, d := range directions {
			row, col := r+d.row, c+d.col

			dfs(row, col, visit, heights[r][c])
		}
	}

	for c := range cols {
		dfs(0, c, pac, heights[0][c])
		dfs(rows-1, c, atl, heights[rows-1][c])
	}

	for r := range rows {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, cols-1, atl, heights[r][cols-1])
	}

	for r := range rows {
		for c := range cols {
			pos := position{r, c}

			if pac[pos] && atl[pos] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}
