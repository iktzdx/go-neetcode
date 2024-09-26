package searcha2dmatrix

func SearchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	if target > matrix[rows-1][cols-1] {
		return false
	}

	for i := range rows {
		if matrix[i][0] <= target && target <= matrix[i][cols-1] {
			return search(matrix[i], target)
		}
	}

	return false
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
