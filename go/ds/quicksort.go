package ds

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	var less []int
	var greater []int

	for _, num := range nums[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	lS := quicksort(less)
	gS := quicksort(greater)
	res := append(lS, pivot)
	res = append(res, gS...)
	return res
}
