package kata

import "slices"

// ::KATA START::
func threeSum(nums []int) [][]int {
	results := make([][]int, 0, len(nums))

	for i, n := 0, len(nums); i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if !containsSlice(results, []int{nums[i], nums[j], nums[k]}) {
						results = append(results, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}

	}

	return results
}

func containsSlice(arrays [][]int, target []int) bool {
	for _, arr := range arrays {
		if slices.Equal(arr, target) {
			return true
		}
	}
	return false
}

// ::KATA END::
