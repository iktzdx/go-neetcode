package maxareaofisland

type dfsFunc func(r, c int) int

func MaxAreaOfIsland(grid [][]int) int {
	var (
		area int
		dfs  dfsFunc
	)

	rows, cols := len(grid), len(grid[0])

	dfs = func(r, c int) int {
		if r < 0 || r == rows || c < 0 || c == cols || grid[r][c] == 0 || grid[r][c] == -1 {
			return 0
		}

		grid[r][c] = -1

		return (1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1))
	}

	for r := range rows {
		for c := range cols {
			area = max(area, dfs(r, c))
		}
	}

	return area
}
