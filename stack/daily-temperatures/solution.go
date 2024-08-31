package dailytemperatures

func DailyTemperatures(temperatures []int) []int {
	if len(temperatures) < 2 {
		return []int{0}
	}

	answer := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				answer[i] = j - i
				break
			}
		}
	}

	return answer
}
