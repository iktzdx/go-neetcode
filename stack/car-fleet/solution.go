package carfleet

func CarFleet(target int, position []int, speed []int) int {
	if len(position) < 2 {
		return 1
	}

	times := make([]float32, target)
	for idx, pos := range position {
		times[pos] = float32(target-pos) / float32(speed[idx])
	}

	var fleets int
	var fastest float32
	for i := target - 1; i >= 0; i-- {
		if times[i] > fastest {
			fastest = times[i]
			fleets++
		}
	}

	return fleets
}
