package dailytemperatures

func DailyTemperatures(temperatures []int) []int {
	if len(temperatures) < 2 {
		return []int{0}
	}

	stack := make([]int, 0)
	answer := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			answer[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return answer
}
