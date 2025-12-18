package kata

// ::KATA START::
func majorityElement(nums []int) int {
	majority := 0
	count := map[int]int{}

	for _, num := range nums {
		count[num]++
		if count[num] > count[majority] {
			majority = num
		}
	}
	return majority
}

// ::KATA END::
