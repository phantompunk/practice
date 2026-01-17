package kata

// ::KATA START::
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	subtotal := 0
	maximum := nums[0]
	for _, num := range nums {
		subtotal += num
		maximum = max(subtotal, maximum)
		if subtotal < 0 {
			subtotal = 0
		}
	}
	return maximum

	// brute force
	// currMax := nums[0]
	// for i, n := 0, len(nums); i < n; i++ {
	//    currMax = max(nums[i], currMax)
	//
	//    sum := nums[i]
	// 	for j := i + 1; j < n; j++ {
	//      sum += nums[j]
	//      currMax = max(sum, currMax)
	// 	}
	// }
	//  return currMax
}

// ::KATA END::
