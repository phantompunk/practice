package kata

// ::KATA START::
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit := 0
	for i, buy := range prices {
		for j := i + 1; j < len(prices); j++ {
      sell := prices[j]
			maxProfit = max(sell - buy, maxProfit)
		}
	}
	return maxProfit
}

// ::KATA END::
