package kata

// ::KATA START::
func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	hashSet := map[int]int{}

	for idx, num := range nums {
		diff := target - num
		if val, ok := hashSet[num]; ok {
			result[0] = val
			result[1] = idx
			break
		}
		hashSet[diff] = idx
	}

	return result
}

// ::KATA END::
