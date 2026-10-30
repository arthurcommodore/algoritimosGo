package algoritimo

func Binary_search(num int, nums []int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == num {
			return mid
		} else if nums[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
