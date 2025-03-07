package kata

func majorityElement(nums []int) int {
	majority := 0
	hashMap := map[int]int{}

	for _, num := range nums {
		hashMap[num]++
		if hashMap[num] > hashMap[majority] {
			majority = num
		}
	}

	return majority
}
