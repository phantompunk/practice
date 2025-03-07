package kata

func containsDuplicate(nums []int) bool {
	freq := map[int]int{}
	for _, num := range nums {
		if _, exists := freq[num]; exists {
			return true
		}
		freq[num]++
	}

	return false
}
