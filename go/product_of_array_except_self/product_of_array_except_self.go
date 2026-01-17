package kata

// ::KATA START::
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	prefix := 1
	for i := range nums {
		answer[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= postfix
		postfix *= nums[i]
	}

	return answer
}

// ::KATA END::

// ans := make([]int, len(nums))
// for i := range nums {
// 	sub := 1
// 	for j := range nums {
// 		if i == j {
// 			continue
// 		}
// 		sub *= nums[j]
// 	}
// 	ans[i] = sub
// }
// return ans
