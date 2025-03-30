package kata

func containsDuplicate(nums []int) bool {
	seen := map[int]struct{}{}
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
