package wallsandgates

import "math"

const (
	Gate  = 0
	Wall  = -1
	Empty = math.MaxInt
)

type position struct {
	row, col int
}

func WallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}

	rows, cols := len(rooms), len(rooms[0])

	q := make([]position, 0)

	// Go through all rooms and only add the gates to the queue.
	for r := range rows {
		for c := range cols {
			if rooms[r][c] == Gate {
				q = append(q, position{r, c})
			}
		}
	}

	directions := []position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// While we have more nodes to process ...
	for len(q) > 0 {
		// ... dequeue a node.
		pos := q[0]
		q = q[1:]

		// Look in every direction for an empty room: up, down, left, right.
		for _, d := range directions {
			row, col := pos.row+d.row, pos.col+d.col

			// if the node we are trying to look at is out of bounds or is a wall, give up.
			if (row < 0 || row >= rows) || (col < 0 || col >= cols) || rooms[row][col] != Empty {
				continue
			}

			// Since the node is a valid empty room, mark it with the current distance from the gate.
			rooms[row][col] = rooms[pos.row][pos.col] + 1

			// Also, add it to the queue so we can process it and continue the BFS.
			q = append(q, position{row, col})
		}
	}
}
