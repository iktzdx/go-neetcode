package searcha2dmatrix

func SearchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	top, bottom, curr := 0, rows-1, 0

	if target > matrix[bottom][cols-1] {
		return false
	}

	for top <= bottom {
		curr = top + ((bottom - top) / 2)
		row := matrix[curr]

		if target == row[0] || target == row[cols-1] {
			return true
		}

		if target > row[cols-1] {
			top = curr + 1
		} else if target < row[0] {
			bottom = curr - 1
		} else {
			break
		}
	}

	return search(matrix[curr], target)
}

func search(row []int, target int) bool {
	leftIdx, rightIdx := 0, len(row)-1

	for leftIdx <= rightIdx {
		midIdx := leftIdx + ((rightIdx - leftIdx) / 2)
		currValue := row[midIdx]

		if target == currValue {
			return true
		}

		if target > currValue {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}

	return false
}
