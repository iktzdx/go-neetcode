package twosumiiinputarrayissorted

func TwoSum(numbers []int, target int) []int {
	length := len(numbers)

	if length == 2 {
		return []int{1, 2}
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}

			if numbers[i]+numbers[j] > target {
				length = j + 1
			}
		}
	}

	return nil
}
