package dailytemperatures

type stack [][2]int

func DailyTemperatures(temperatures []int) []int {
	if len(temperatures) < 2 {
		return []int{0}
	}

	var s stack
	answer := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for !s.isEmpty() && s.top()[0] < temp {
			answer[s.top()[1]] = i - s.top()[1]
			s.pop()
		}

		s.push([2]int{temp, i})
	}

	return answer
}

func (s *stack) top() [2]int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() {
	(*s) = (*s)[:len(*s)-1]
}

func (s *stack) push(v [2]int) {
	(*s) = append(*s, v)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}
