package algoritimo

func Selection_sort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		sort(nums, i)
	}

	return nums
}

func sort(nums []int, indice int) {
	var menorIndice int = indice

	for i := indice + 1; i < len(nums); i++ {
		if nums[menorIndice] > nums[i] {
			menorIndice = i
		}
	}

	nums[indice], nums[menorIndice] = nums[menorIndice], nums[indice]
}
