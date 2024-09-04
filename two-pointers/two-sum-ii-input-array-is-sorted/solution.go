package twosumiiinputarrayissorted

func TwoSum(numbers []int, target int) []int {
	length := len(numbers)

	if length == 2 {
		return []int{1, 2}
	}

	left, right := 0, length-1
	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
			continue
		}

		if sum < target {
			left++
			continue
		}

		break
	}

	return []int{left + 1, right + 1}
}
