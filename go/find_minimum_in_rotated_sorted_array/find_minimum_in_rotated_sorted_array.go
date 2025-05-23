package kata

func findMin(nums []int) int {
	lo, mid, hi := 0, 0, len(nums)-1

	for lo < hi {
		mid = lo + (hi-lo)/2

		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
