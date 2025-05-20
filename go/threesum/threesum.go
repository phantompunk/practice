package kata

import "sort"

func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)

	// iterate thru values
	for i, n := range nums {
		// i is the start or the same as the previous continue
		if i > 0 && n == nums[i-1] {
			continue
		}

		// set your pointers
		l, r := i+1, len(nums)-1
		// loop thru while l<r
		for l < r {
			// calculate the sume
			x := n + nums[l] + nums[r]
			if x > 0 {
				r--
			} else if x < 0 {
				l++
			} else {
				results = append(results, []int{n, nums[l], nums[r]})
				l++
				// skip where the values are duplicated
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return results
}
