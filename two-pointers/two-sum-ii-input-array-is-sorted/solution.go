package twosumiiinputarrayissorted

func TwoSum(numbers []int, target int) []int {
	length := len(numbers)

	if length == 2 {
		return []int{1, 2}
	}

	left, right := 0, length-1
	for numbers[left]+numbers[right] != target {
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}
