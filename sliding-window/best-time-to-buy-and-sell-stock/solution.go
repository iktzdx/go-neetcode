package besttimetobuyandsellstock

func MaxProfit(prices []int) int {
	var maxProfit int

	left, right := 0, 1
	for right < len(prices) {
		if prices[left] < prices[right] {
			maxProfit = max(maxProfit, prices[right]-prices[left])
		} else {
			left = right
		}

		right++
	}

	return maxProfit
}
