package kthlargestelementinanarray

type quickSelectFunc func(l, r int) int

func FindKthLargest(nums []int, k int) int {
	k = len(nums) - k

	var qs quickSelectFunc

	qs = func(l, r int) int {
		pivot, p := nums[r], l

		for i := l; i < r; i++ {
			if nums[i] <= pivot {
				nums[p], nums[i] = nums[i], nums[p]
				p++
			}
		}

		nums[p], nums[r] = nums[r], nums[p]

		if p > k {
			return qs(l, p-1)
		} else if p < k {
			return qs(p+1, r)
		} else {
			return nums[p]
		}
	}

	return qs(0, len(nums)-1)
}
