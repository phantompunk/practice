package kata

func minEatingSpeed(piles []int, h int) int {
	lo, mid, hi := 0, 0, maxPile(piles)

	for lo < hi {
		mid = lo + (hi-lo)/2
		if canFinishPile(piles, mid, h) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canFinishPile(piles []int, k, h int) bool {
	count := 0
	for _, pile := range piles {
		count += pile / k
		if pile%k != 0 {
			count++
		}
	}
	return count <= h
}

func maxPile(piles []int) int {
	maxP := 0
	for _, pile := range piles {
		maxP = max(maxP, pile)
	}
	return maxP
}
