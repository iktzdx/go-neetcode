package carfleet

func CarFleet(target int, position []int, speed []int) int {
	var (
		fleets  int
		fastest float32
	)

	if len(position) < 2 {
		return 1
	}

	times := make([]float32, target)
	for idx, pos := range position {
		times[pos] = float32(target-pos) / float32(speed[idx])
	}

	for i := target - 1; i >= 0; i-- {
		if times[i] > fastest {
			fastest = times[i]
			fleets++
		}
	}

	return fleets
}
