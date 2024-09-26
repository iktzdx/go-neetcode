package searcha2dmatrix

func SearchMatrix(matrix [][]int, target int) bool {
	var result bool

	leftIdx, rightIdx := 0, len(matrix)-1

	if target > matrix[rightIdx][len(matrix[rightIdx])-1] {
		return result
	}

	for leftIdx <= rightIdx {
		midIdx := leftIdx + ((rightIdx - leftIdx) / 2)
		midRow := matrix[midIdx]

		if target == midRow[0] || target == midRow[len(midRow)-1] {
			return true
		}

		if target > midRow[0] && target < midRow[len(midRow)-1] {
			result = search(midRow, target)
		}

		if target > midRow[0] && target > midRow[len(midRow)-1] {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}

	}

	return result
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
