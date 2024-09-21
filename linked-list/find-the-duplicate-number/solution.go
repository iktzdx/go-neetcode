package findtheduplicatenumber

func FindDuplicate(nums []int) int {
	var slow, fast int

	// Find the intersection point.
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Find the entrance to the cycle.
	var slow2 int
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}

	return slow
}
