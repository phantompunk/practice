package kata

// ::KATA START::
func containsDuplicate(nums []int) bool {
	hashSet := map[int]struct{}{}
	for _, num := range nums {
		if _, exists := hashSet[num]; exists {
			return true
		}
		hashSet[num] = struct{}{}
	}

	return false 
}
// ::KATA END::
