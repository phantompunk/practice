package kata

// ::KATA START::
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		diff := numbers[i] + numbers[j]
		if diff < target {
			i++
		} else if diff > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

// ::KATA END::
