package kata

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

//::LEETCODE::
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for idx, num := range nums {
		diff := target - num
		if val, exists := m[diff]; exists {
			return []int{val, idx}
		}
		m[num] = idx
	}
	return nil
}
//::LEETCODE::
