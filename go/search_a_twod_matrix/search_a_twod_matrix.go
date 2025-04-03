package kata

func searchMatrix(matrix [][]int, target int) bool {
	end := len(matrix[0]) - 1

	for _, row := range matrix {
		if target <= row[end] {
			return search(row, target)
		}
	}
	return false
}

func search(nums []int, target int) bool {
	lo, mid, hi := 0, 0, len(nums)-1

	for lo <= hi {
		mid = (lo + hi) / 2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			return true
		}
	}
	return false
}
