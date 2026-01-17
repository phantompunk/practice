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

// Brute force solution: Time O(n^2) Space O(1)
// func majorityElement(nums []int) int {
// 	n := len(nums)
//
// 	for i := range n {
// 		count := 0
// 		for j := range n {
// 			if nums[i] == nums[j] {
// 				count++
// 			}
// 		}
//
// 		if count > n/2 {
// 			return nums[i]
// 		}
// 	}
//
// 	return -1
// }
