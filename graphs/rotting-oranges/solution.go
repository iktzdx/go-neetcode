package rottingoranges

const (
	emptyCell    = 0
	freshOrange  = 1
	rottenOrange = 2
)

type position struct {
	row, col int
}

func OrangesRotting(grid [][]int) int {
	var fresh, time int

	rows, cols := len(grid), len(grid[0])
	queue := make([]position, 0)

	for r := range rows {
		for c := range cols {
			if grid[r][c] == freshOrange {
				fresh++
			}

			if grid[r][c] == rottenOrange {
				queue = append(queue, position{r, c})
			}
		}
	}

	directions := []position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for fresh > 0 && len(queue) > 0 {
		length := len(queue)

		for range length {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range directions {
				row, col := curr.row+d.row, curr.col+d.col

				if row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] == freshOrange {
					grid[row][col] = rottenOrange

					queue = append(queue, position{row, col})

					fresh--
				}
			}
		}

		time++
	}

	if fresh == 0 {
		return time
	}

	return -1
}
