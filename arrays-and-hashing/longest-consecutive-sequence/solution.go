package longestconsecutivesequence

func LongestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	longest := 1
	numsMap := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		numsMap[n] = struct{}{}
	}

	for _, n := range nums {
		if _, ok := numsMap[n-1]; !ok { // check if it's the start of the sequence
			length := 0
			for length < len(nums) {
				if _, ok := numsMap[n+length]; !ok { // search for the end of the sequence
					break
				}
				length++
			}
			longest = max(longest, length)
		}
	}

	return longest
}
