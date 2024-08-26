package longestconsecutivesequence

func LongestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	longest := 1
	numsMap := make(map[int]bool, len(nums))

	for _, n := range nums {
		numsMap[n] = true
	}

	for _, n := range nums {
		if numsMap[n+1] { // search for the end of the sequence
			continue
		}

		length := 0
		for curr := n; numsMap[curr]; curr-- { // check if it's the start of the sequence
			length++
		}

		longest = max(longest, length)
	}

	return longest
}
