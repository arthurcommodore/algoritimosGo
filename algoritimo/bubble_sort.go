package algoritimo

func Bubble_sort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {

			atual := nums[j]
			prox := nums[j+1]
			if atual > prox {
				nums[j] = prox
				nums[j+1] = atual
			}
		}
	}
	return nums
}
